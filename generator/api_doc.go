package generator

import (
	"net/url"
	"regexp"
	"strings"
	"sync"

	"github.com/PuerkitoBio/goquery"
	"github.com/sirupsen/logrus"

	"github.com/go-courier/codegen"
	"github.com/davecgh/go-spew/spew"
)

type APIDocEntries struct {
	ContentsURL  string
	DataTypesURL string
}

func NewAPIDoc(apiDocEntries *APIDocEntries) *APIDoc {
	return &APIDoc{
		APIDocEntries: apiDocEntries,
		APIs:          map[string]*API{},
	}
}

type APIDoc struct {
	*APIDocEntries
	Service   string
	APIs      map[string]*API
	DataTypes map[string]*TypeSchema
}

func (doc *APIDoc) Scan() {
	doc.scanAPIContents()
	doc.scanAPIDetails()
	doc.scanDataTypes()
}

func (doc *APIDoc) addDataType(s *TypeSchema) {
	if doc.DataTypes == nil {
		doc.DataTypes = map[string]*TypeSchema{}
	}
	doc.DataTypes[s.Name] = s
}

func (doc *APIDoc) scanDataTypes() {
	d, err := goquery.NewDocument(doc.DataTypesURL)
	if err != nil {
		logrus.Panicf("%s", err)
	}

	d.Find("#docArticleContent > h2[id]").Each(func(i int, h2 *goquery.Selection) {
		ts := &TypeSchema{}
		ts.Name = h2.Text()

		doc.addDataType(ts)

		p := h2.Next()
		ts.Desc = p.Text()

		p.Next().Find("tbody tr").Each(func(i int, tr *goquery.Selection) {
			cols := tr.Find("td")

			propType := &TypeSchema{}

			col := cols.Eq(1)

			if col.Find("a").Length() != 0 {
				propType.Name = col.Find("a").First().Text()
			} else {
				propType.Type = IndirectType(col.Text())
			}

			if IsArrayType(col.Text()) {
				propType = &TypeSchema{
					Items: propType,
				}
			}

			propType.Desc = cols.Eq(3).Text()
			propType.Required = cols.Eq(2).Text() != "否"

			ts.AddProp(cols.Eq(0).Text(), propType)
		})
	})
}

func (doc *APIDoc) scanAPIContents() {
	d, err := goquery.NewDocument(doc.ContentsURL)
	if err != nil {
		logrus.Panicf("%s", err)
	}

	originURL, _ := url.Parse(doc.ContentsURL)

	d.Find("#docArticleContent tbody tr").Each(func(i int, s *goquery.Selection) {
		cols := s.Find("td")

		name := cols.Eq(0).Text()
		docURL, _ := cols.Find("a").Eq(0).Attr("href")

		u, _ := url.Parse(docURL)
		u = originURL.ResolveReference(u)

		desc := cols.Eq(1).Text()

		doc.APIs[name] = NewAPI(&doc.Service, name, desc, u.String())
	})
}

func (doc *APIDoc) scanAPIDetails() {
	wg := sync.WaitGroup{}
	wg.Add(len(doc.APIs))

	jobs := make(chan func())

	for i := 0; i < 10; i++ {
		go func() {
			for runJob := range jobs {
				runJob()
				wg.Add(-1)
			}
		}()
	}

	for name := range doc.APIs {
		api := doc.APIs[name]

		if doc.Service == "" {
			api.ScanServiceName()
		}

		jobs <- func() {
			api.Scan()
		}
	}

	wg.Wait()
	close(jobs)
}

func (doc *APIDoc) WriteAPIs() {
	for name := range doc.APIs {
		api := doc.APIs[name]
		pkgName := codegen.LowerSnakeCase(doc.Service)
		f := codegen.NewFile(pkgName, "./clients/"+pkgName+"/"+codegen.LowerSnakeCase(api.Name)+".go")

		api.Write(f)

		f.WriteFile()
		logrus.Infof("Generated %s.%s", doc.Service, api.Name)
	}
}

func (doc *APIDoc) WriteDataTypes() {
	pkgName := codegen.LowerSnakeCase(doc.Service)
	f := codegen.NewFile(pkgName, "./clients/"+pkgName+"/data_types.go")

	for name := range doc.DataTypes {
		dataType := doc.DataTypes[name]
		dataType.Write(f)
	}

	_, err := f.WriteFile()
	if err != nil {
		panic(err)
	}
	logrus.Infof("Generated data types of %s", doc.Service)
}

func (doc *APIDoc) SyncAndWriteFiles() {
	doc.WriteDataTypes()
	doc.WriteAPIs()
}

func NewAPI(service *string, name string, desc string, docUrl string) *API {
	regionType := BasicTypeString

	return &API{
		Service: service,
		Name:    name,
		Desc:    desc,
		DocURL:  docUrl,
		Parameters: &TypeSchema{
			Name: codegen.UpperCamelCase(name + "Request"),
			Desc: desc + "\n" + docUrl,
			Tag:  "name",
			Props: map[string]*TypeSchema{
				"Region": &TypeSchema{
					Type:     &regionType,
					Desc:     "区域",
					Required: true,
				},
			},
		},
		Response: &TypeSchema{
			Name: codegen.UpperCamelCase(name + "Response"),
			AllOf: []*TypeSchema{
				&TypeSchema{
					ImportPath: "github.com/morlay/goqcloud",
					Name:       "TencentCloudBaseResponse",
					Required:   true,
				},
			},
			Props: map[string]*TypeSchema{},
		},
	}
}

type API struct {
	Service    *string
	Name       string
	Version    string
	Desc       string
	DocURL     string
	Parameters *TypeSchema
	Response   *TypeSchema
}

func (api *API) Write(file *codegen.File) {
	api.Parameters.Write(file)
	api.WriteInvoke(file)
	api.Response.Write(file)
}

func (api *API) WriteInvoke(file *codegen.File) {

	spew.Dump(file.Use("github.com/morlay/goqcloud", "Client"))

	file.WriteBlock(
		codegen.Func(
			codegen.Var(codegen.Type(file.Use("github.com/morlay/goqcloud", "Client")), "client"),
		).Return(
			codegen.Var(codegen.Star(codegen.Type(api.Response.Name))),
			codegen.Var(codegen.Star(codegen.Error)),
		).MethodOf(
			codegen.Var(codegen.Star(codegen.Type(api.Parameters.Name)), "req"),
		).Named("Invoke").Do(
			codegen.Define(codegen.Id("resp")).By(file.Expr("&?{}", codegen.Type(api.Response.Name))),
			codegen.Define(codegen.Id("err")).By(
				codegen.Sel(
					codegen.Id("client"),
					codegen.Call("Request", file.Val(*api.Service), file.Val(api.Name), file.Val(api.Version)),
					codegen.Call("Do", codegen.Id("req"), codegen.Id("resp")),
				),
			),
			codegen.Return(codegen.Id("resp"), codegen.Id("err")),
		),
	)
}

func (api *API) AddParameter(name string, tpe *TypeSchema) {
	name = strings.TrimSuffix(name, ".N")
	api.Parameters.AddProp(name, tpe)
}

func (api *API) AddResponseProp(name string, tpe *TypeSchema) {
	api.Response.AddProp(name, tpe)
}

var reVersion = regexp.MustCompile("[0-9]{4}-[0-9]{2}-[0-9]{2}")
var reTencentCloudHostname = regexp.MustCompile("([a-z]+)\\.tencentcloudapi\\.com")

func (api *API) ScanServiceName() {
	d, err := goquery.NewDocument(api.DocURL)
	if err != nil {
		logrus.Panicf("%s", err)
	}

	service := reTencentCloudHostname.FindStringSubmatch(d.Find("#docArticleContent h2").First().Next().Text())[1]
	*api.Service = service
}

func (api *API) Scan() {
	d, err := goquery.NewDocument(api.DocURL)
	if err != nil {
		logrus.Panicf("%s", err)
	}

	tables := make([][]*goquery.Selection, 3)
	tableIndex := 0

	logrus.Infof("Scan api %s.%s from `%s`", *api.Service, api.Name, api.DocURL)

	d.Find("#docArticleContent h2").Siblings().Each(func(i int, s *goquery.Selection) {
		if s.Is("h2") {
			if strings.HasPrefix(s.Text(), "2") {
				tables[tableIndex] = make([]*goquery.Selection, 0)
			}
			if strings.HasPrefix(s.Text(), "3") {
				tableIndex = 1
				tables[tableIndex] = make([]*goquery.Selection, 0)
			}
			if strings.HasPrefix(s.Text(), "4") {
				tableIndex = 2
				tables[tableIndex] = make([]*goquery.Selection, 0)
			}
		}

		if s.Is("table") {
			tables[tableIndex] = append(tables[tableIndex], s)
		}
	})

	for _, table := range tables[0] {
		table.Find("tbody tr").Each(func(i int, s *goquery.Selection) {
			cols := s.Find("td")

			paramType := &TypeSchema{}

			tpeCol := cols.Eq(2)

			if tpeCol.Find("a").Length() != 0 {
				paramType.Name = tpeCol.Find("a").First().Text()
			} else {
				paramType.Type = IndirectType(tpeCol.Text())
			}

			if IsArrayType(tpeCol.Text()) {
				paramType = &TypeSchema{
					Items: paramType,
				}
			}

			paramType.Desc = cols.Eq(3).Text()
			paramType.Required = cols.Eq(1).Text() != "否"

			name := cols.Eq(0).Text()

			switch name {
			case "Action":
			case "Version":
				api.Version = reVersion.FindString(paramType.Desc)
			default:
				api.AddParameter(cols.Eq(0).Text(), paramType)
			}
		})
	}

	for _, table := range tables[1] {
		table.Find("tbody tr").Each(func(i int, s *goquery.Selection) {
			cols := s.Find("td")

			respPropType := &TypeSchema{}

			tpeCol := cols.Eq(1)

			if tpeCol.Find("a").Length() != 0 {
				respPropType.Name = tpeCol.Find("a").First().Text()
			} else {
				respPropType.Type = IndirectType(tpeCol.Text())
			}

			if IsArrayType(tpeCol.Text()) {
				respPropType = &TypeSchema{
					Items: respPropType,
				}
			}

			respPropType.Desc = cols.Eq(2).Text()
			respPropType.Required = true

			name := cols.Eq(0).Text()

			switch name {
			case "RequestId":
			default:
				api.AddResponseProp(name, respPropType)
			}

		})
	}
}
