package controllers

//权限管理 垂直权限  水平权限
//垂直权限: 不同类型的用户权限不一样,例如:admin , user , develop等
//水平权限: 相同类型的用户权限不一样,例如:同为 user ,但是用户之间的数据不能相互展示

//垂直权限设置
//后端设置一个menus,存放在数据库中
//[]module{
//Moudule("user","user_action","用户管理",[controller.action])
//Moudule("task","task_action","任务管理",[controller.action])
//}

//Prepare()
//loginUser.Role => [controller.action]
//c.GetControllerAndAction

//从prepare使用c.GetControllerAndAction获取controller和action
//判断用户的[controller.action]中是否有controller和action来决定用户是否有权限,然后展示不同页面

//水平权限设置
//即在处理数据时加上限制,例如: select * from tasks where userName="siri"
