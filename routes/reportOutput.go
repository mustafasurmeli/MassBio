package routes

import (
	"WebAPI1/database"
	"WebAPI1/models"
	"github.com/gofiber/fiber/v2"
)

type ReportOutput struct {
	MainUploadedVariation string  `json:"main_uploaded_variation"`
	MainExistingVariation string  `json:"main_existing_variation"`
	MainSymbol            string  `json:"main_symbol"`
	MainAfCf              float64 `json:"main_af_cf"`
	MainDp                float64 `json:"main_dp"`
	Details2Provean       string  `json:"details_2_provean"`
	Details2DannScore     float64 `json:"details_2_dann_score"`
	LinksMondo            string  `json:"links_mondo"`
	LinksPhenoPubmed      string  `json:"links_pheno_pubmed"`
}

func CreateResponse(reportOutput models.ReportOutput) ReportOutput {
	return ReportOutput{MainUploadedVariation: reportOutput.MainUploadedVariation, MainExistingVariation: reportOutput.MainExistingVariation,
		MainSymbol: reportOutput.MainSymbol, MainAfCf: reportOutput.MainAfCf, MainDp: reportOutput.MainDp,
		Details2DannScore: reportOutput.Details2DannScore, Details2Provean: reportOutput.Details2Provean,
		LinksMondo: reportOutput.LinksMondo, LinksPhenoPubmed: reportOutput.LinksPhenoPubmed}
}
func GetReportOutputs(c *fiber.Ctx) error {
	reportOutputs := []models.ReportOutput{}
	database.Database.Db.Find(&reportOutputs)
	response := []ReportOutput{}
	for _, reportOutput := range reportOutputs {
		responseRO := CreateResponse(reportOutput)
		response = append(response, responseRO)
	}
	return c.Status(200).JSON(response)
}
