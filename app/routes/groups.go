package routes

import (
	"cbam_api/data"
)

func QuarterlyReportGroup(router *Router, client *data.MongoDBClient) {
	router.Handle("QuarterlyReports", "GET", "/reports", GetAllQuarterlyReports(client))
	router.Handle("QuarterlyReports", "GET", "/reports/{id}", GetQuarterlyReport(client))
	// router.Handle("QuarterlyReports", "POST", "/reports/create", CreateQuarterlyReport(client))
	router.Handle("QuarterlyReports", "DELETE", "/reports/delete/{id}", DeleteAllQuarterlyReports(client))
}

func DeclarantGroup(router *Router, client *data.MongoDBClient) {
	router.Handle("Declarant", "GET", "/reports/{id}/declarent", GetGoodsEmissions(client))
}

func ImporterGroup(router *Router, client *data.MongoDBClient) {
	router.Handle("Importer", "GET", "/reports/{id}/importer", GetImporter(client))
}
func ImportedGoodsGroup(router *Router, client *data.MongoDBClient) {
	router.Handle("ImportedGoods", "GET", "/reports/{id}/imported-goods", GetImportedGoods(client))
}
func GoodsEmissionsGroup(router *Router, client *data.MongoDBClient) {
	router.Handle("GoodsEmissions", "GET", "/reports/{id}/imported-goods/emissions", GetGoodsEmissions(client))
}

func SignaturesGroup(router *Router, client *data.MongoDBClient) {
	router.Handle("Signatures", "GET", "/reports/{id}/signatures", GetSignatures(client))
}

func RemarksEmissionsGroup(router *Router, client *data.MongoDBClient) {
	router.Handle("RemarksEmissions", "GET", "/reports/{id}/imported-goods/{goodId}/remarks-emissions", GetRemarksEmissions(client))
}

func NationalCompetentAuthGroup(router *Router, client *data.MongoDBClient) {
	router.Handle("NationalCompetentAuth", "GET", "/reports/{id}/national-competent-auth", GetNationalCompetentAuth(client))
}

func SupportingDocumentsGroup(router *Router, client *data.MongoDBClient) {
	router.Handle("Documents", "GET", "/reports/{id}/imported-goods/{goodId}/supporting-documents", GetSupportingDocuments(client))
}

func FileGroup(router *Router, client *data.MongoDBClient) {
	router.Handle("Files", "POST", "/api/upload", UploadHandler(client))

}
