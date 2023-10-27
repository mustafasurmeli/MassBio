package models

type ReportOutput struct {
	Row                   int64   `json:"row" gorm:"primaryKKey"`
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
