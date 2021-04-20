package errcode

var (
	ErrorGetTagListFail = NewError(20010001, "get tag list failed")
	ErrorCreateTagFail  = NewError(20010002, "create tag failed")
	ErrorUpdateTagFail  = NewError(20010003, "update tag failed")
	ErrorDeleteTagFail  = NewError(20010004, "delete tag failed")
	ErrorCountTagFail   = NewError(20010005, "count tag failed")

	ErrorGetArticleFail    = NewError(20020001, "get article failed")
	ErrorGetArticlesFail   = NewError(20020002, "get articles failed")
	ErrorCreateArticleFail = NewError(20020003, "create article failed")
	ErrorUpdateArticleFail = NewError(20020004, "update article failed")
	ErrorDeleteArticleFail = NewError(20020005, "delete article failed")

	ErrorUploadFileFail = NewError(20030001, "upload file failed")
)
