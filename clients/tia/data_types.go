package tia

// 训练任务信息

type Job *Job

// 日志

type Log *Log

// 用于描述模型的详细情况
// "Model": {
// "Name": "test-model",
// "Description": "test-model",
// "Cluster": "ap-beijing",
// "Model": "cos://test-1255502019.cos.ap-shanghai.myqcloud.com/example:/data/mnist",
// "RuntimeVersion": "tiaserv-1.6.0-cpu",
// "CreateTime": "2018-04-26 15:59:25 +0800 CST",
// "State": "Running",
// "ServingUrl": "140.143.51.230",
// "Message": "Deployment does not have minimum availability.",
// "AppId": 1255502019,
// "ServType": "1U2G0P"
// },

type Model *Model
