package routes

func QuarterlyReportGroup(router *Router) {
	router.Handle("QuarterlyReports", "GET", "/reports/mockreport", ExampleReportHandler())
}
