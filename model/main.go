package model

import (
	"encoding/xml"

	"github.com/shopspring/decimal"
)

// swagger:model QReport
type QReport struct {
	// XMLName               xml.Name              `xml:"QReport"`
	SubmissionDate        SimpleDate            `xml:"SubmissionDate"`
	ReportId              string                `xml:"ReportId"`
	ReportingPeriod       string                `xml:"ReportingPeriod"`
	Year                  string                `xml:"Year"`
	Declarant             Declarant             `xml:"Declarant"`
	Importer              Importer              `xml:"Importer"`
	NationalCompetentAuth NationalCompetentAuth `xml:"NationalCompetentAuth"`
	Signatures            Signatures            `xml:"Signatures"`
	ImportedGoods         ImportedGoods         `xml:"ImportedGood"`
}

// swagger:model Declarant
type Declarant struct {
	IdentificationNumber string  `xml:"IdentificationNumber"`
	Name                 string  `xml:"Name"`
	Role                 string  `xml:"Role"`
	ActorAddress         Address `xml:"ActorAddress"`
}

// swagger:model Importer
type Importer struct {
	IdentificationNumber string  `xml:"IdentificationNumber"`
	Name                 string  `xml:"Name"`
	ImporterAddress      Address `xml:"ImporterAddress"`
}

// swagger:model Address
type Address struct {
	Country              string          `xml:"Country"`
	SubDivision          string          `xml:"SubDivision"`
	City                 string          `xml:"City"`
	Street               string          `xml:"Street"`
	StreetAdditionalLine string          `xml:"StreetAdditionalLine"`
	Number               decimal.Decimal `xml:"Number"`
	Postcode             string          `xml:"Postcode"`
	POBox                string          `xml:"POBox"`
}

// swagger:model NationalCompetentAuth
type NationalCompetentAuth struct {
	ReferenceNumber string `xml:"ReferenceNumber"`
}

// swagger:model Signatures
type Signatures struct {
	ReportConfirmation               ReportConfirmation                `xml:"ReportConfirmation"`
	ApplicableMethdologyConfirmation ApplicableMethodologyConfirmation `xml:"ApplicableMethdologyConfirmation"`
}

// swagger:model ReportConfirmation
type ReportConfirmation struct {
	GlobalDataConfirmation  bool   `xml:"GlobalDataConfirmation"`
	UseOfDataConfirmation   bool   `xml:"UseOfDataConfirmation"`
	SignatureDate           string `xml:"SignatureDate"`
	SignaturePlace          string `xml:"SignaturePlace"`
	Signature               string `xml:"Signature"`
	PositionOfPersonSending string `xml:"PositionOfPersonSending"`
}

// swagger:model ApplicableMethodologyConfirmation
type ApplicableMethodologyConfirmation struct {
	OtherApplicableReportingMethodology bool `xml:"OtherApplicableReportingMethodology"`
}

// swagger:model ImportedGoods
type ImportedGoods struct {
	ItemNumber          string             `xml:"ItemNumber"`
	CommodityCode       CommodityCode      `xml:"CommodityCode"`
	OriginCountry       OriginCountry      `xml:"OriginCountry"`
	ImportedQuantity    ImportedQuantity   `xml:"ImportedQuantity"`
	MeasureImported     MeasureImported    `xml:"MeasureImported"`
	TotalEmissions      TotalEmissions     `xml:"TotalEmissions"`
	SupportingDocuments SupportingDocument `xml:"SupportingDocuments"`
	Remarks             Remarks            `xml:"Remarks"`
	GoodsEmissions      []GoodsEmission    `xml:"GoodsEmissions"`
}

// swagger:model CommodityCode
type CommodityCode struct {
	HsCode           string           `xml:"HsCode"`
	CnCode           string           `xml:"CnCode"`
	CommodityDetails CommodityDetails `xml:"CommodityDetails"`
}

// swagger:model CommodityDetails
type CommodityDetails struct {
	Description string `xml:"Description"`
}

// swagger:model OriginCountry
type OriginCountry struct {
	Country string `xml:"Country"`
}

// swagger:model ImportedQuantity
type ImportedQuantity struct {
	SequenceNumber           decimal.Decimal          `xml:"SequenceNumber"`
	Procedure                Procedure                `xml:"Procedure"`
	ImportArea               ImportArea               `xml:"ImportArea"`
	MeasureProcedureImported MeasureProcedureImported `xml:"MeasureProcedureImported"`
	SpecialReferences        SpecialReferences        `xml:"SpecialReferences"`
}

// swagger:model Procedure
type Procedure struct {
	RequestedProc string `xml:"RequestedProc"`
	PreviousProc  string `xml:"PreviousProc"`
}

// swagger:model ImportArea
type ImportArea struct {
	ImportArea string `xml:"ImportArea"`
}

// swagger:model MeasureProcedureImported
type MeasureProcedureImported struct {
	Indicator       string `xml:"Indicator"`
	NetMass         string `xml:"NetMass"`
	MeasurementUnit string `xml:"MeasurementUnit"`
}

// swagger:model SpecialReferences
type SpecialReferences struct {
	AdditionalInformation string `xml:"AdditionalInformation"`
}

// swagger:model MeasureImported
type MeasureImported struct {
	NetMass         string `xml:"NetMass"`
	MeasurementUnit string `xml:"MeasurementUnit"`
}

// swagger:model TotalEmissions
type TotalEmissions struct {
	EmissionsPerUnit string `xml:"EmissionsPerUnit"`
	OverallEmissions string `xml:"OverallEmissions"`
	TotalDirect      string `xml:"TotalDirect"`
	TotalIndirect    string `xml:"TotalIndirect"`
	MeasurementUnit  string `xml:"MeasurementUnit"`
}

// swagger:model SupportingDocument
type SupportingDocument struct {
	SequenceNumber    decimal.Decimal `xml:"SequenceNumber"`
	Type              string          `xml:"Type"`
	Country           string          `xml:"Country"`
	ReferenceNumber   string          `xml:"ReferenceNumber"`
	LineItemNumber    string          `xml:"LineItemNumber"`
	IssuingAuthName   string          `xml:"IssuingAuthName"`
	ValidityStartDate SimpleDate      `xml:"ValidityStartDate"`
	ValidityEndDate   SimpleDate      `xml:"ValidityEndDate"`
	Description       string          `xml:"Description"`
	Attachment        Attachment      `xml:"Attachment"`
}

// swagger:model Attachment
type Attachment struct {
	Filename string `xml:"Filename"`
	MIME     string `xml:"MIME"`
}

// swagger:model Remarks
type Remarks struct {
	AdditionalInformation string `xml:"AdditionalInformation"`
}

// swagger:model GoodsEmission
type GoodsEmission struct {
	SequenceNumber             decimal.Decimal              `xml:"SequenceNumber"`
	ProductionCountry          string                       `xml:"ProductionCountry"`
	InstallationOperator       InstallationOperator         `xml:"InstallationOperator"`
	Installation               Installation                 `xml:"Installation"`
	ProducedMeasure            ProducedMeasure              `xml:"ProducedMeasure"`
	InstallationEmissions      InstallationEmissions        `xml:"InstallationEmissions"`
	DirectEmissions            DirectEmissions              `xml:"DirectEmissions"`
	IndirectEmissions          IndirectEmissions            `xml:"IndirectEmissions"`
	ProdMethodQualifyingParams []ProdMethodQualifyingParams `xml:"ProdMethodQualifyingParams"`
	SupportingDocuments        SupportingDocument           `xml:"SupportingDocuments"`
	RemarksEmissions           RemarksEmissions             `xml:"RemarksEmissions"`
	CarbonPriceDue             CarbonPriceDue               `xml:"CarbonPriceDue"`
}

// swagger:model InstallationOperator
type InstallationOperator struct {
	OperatorId      string         `xml:"OperatorId"`
	OperatorName    string         `xml:"OperatorName"`
	OperatorAddress Address        `xml:"OperatorAddress"`
	ContactDetails  ContactDetails `xml:"ContactDetails"`
}

// swagger:model ContactDetails
type ContactDetails struct {
	Name  string `xml:"Name"`
	Phone string `xml:"Phone"`
	Email string `xml:"Email"`
}

// swagger:model Installation
type Installation struct {
	InstallationId   string              `xml:"InstallationId"`
	InstallationName string              `xml:"InstallationName"`
	Address          InstallationAddress `xml:"Address"`
}

// swagger:model InstallationAddress
type InstallationAddress struct {
	EstablishmentCountry string          `xml:"EstablishmentCountry"`
	SubDivision          string          `xml:"SubDivision"`
	City                 string          `xml:"City"`
	Street               string          `xml:"Street"`
	StreetAdditionalLine string          `xml:"StreetAdditionalLine"`
	Number               decimal.Decimal `xml:"Number"`
	Postcode             string          `xml:"Postcode"`
	POBox                string          `xml:"POBox"`
	PlotParcelNumber     string          `xml:"PlotParcelNumber"`
	Latitude             decimal.Decimal `xml:"Latitude"`
	Longitude            decimal.Decimal `xml:"Longitude"`
	CoordinatesType      string          `xml:"CoordinatesType"`
}

// swagger:model ProducedMeasure
type ProducedMeasure struct {
	NetMass         string `xml:"NetMass"`
	MeasurementUnit string `xml:"MeasurementUnit"`
}

// swagger:model InstallationEmissions
type InstallationEmissions struct {
	OverallEmissions string `xml:"OverallEmissions"`
	TotalDirect      string `xml:"TotalDirect"`
	TotalIndirect    string `xml:"TotalIndirect"`
	MeasurementUnit  string `xml:"MeasurementUnit"`
}

// swagger:model DirectEmissions
type DirectEmissions struct {
	DeterminationType                  string `xml:"DeterminationType"`
	ApplicableReportingTypeMethodology string `xml:"ApplicableReportingTypeMethodology"`
	ApplicableReportingMethodology     string `xml:"ApplicableReportingMethodology"`
	SpecificEmbeddedEmissions          string `xml:"SpecificEmbeddedEmissions"`
	MeasurementUnit                    string `xml:"MeasurementUnit"`
}

// swagger:model IndirectEmissions
type IndirectEmissions struct {
	DeterminationType         string `xml:"DeterminationType"`
	SpecificEmbeddedEmissions string `xml:"SpecificEmbeddedEmissions"`
	MeasurementUnit           string `xml:"MeasurementUnit"`
	ElectricitySource         string `xml:"ElectricitySource"`
	OtherSourceIndication     string `xml:"OtherSourceIndication"`
}

// swagger:model ProdMethodQualifyingParams
type ProdMethodQualifyingParams struct {
	SequenceNumber               decimal.Decimal       `xml:"SequenceNumber"`
	MethodId                     string                `xml:"MethodId"`
	MethodName                   string                `xml:"MethodName"`
	SteelMillNumber              string                `xml:"SteelMillNumber"`
	AdditionalInformation        string                `xml:"AdditionalInformation"`
	DirectQualifyingParameters   []QualifyingParameter `xml:"DirectQualifyingParameters"`
	IndirectQualifyingParameters []QualifyingParameter `xml:"IndirectQualifyingParameters"`
}

// swagger:model QualifyingParameter
type QualifyingParameter struct {
	SequenceNumber        decimal.Decimal `xml:"SequenceNumber"`
	DeterminationType     string          `xml:"DeterminationType"`
	ParameterId           string          `xml:"ParameterId"`
	ParameterName         string          `xml:"ParameterName"`
	Description           string          `xml:"Description"`
	ParameterValueType    string          `xml:"ParameterValueType"`
	ParameterValue        string          `xml:"ParameterValue"`
	AdditionalInformation string          `xml:"AdditionalInformation"`
}

// swagger:model RemarksEmissions
type RemarksEmissions struct {
	SequenceNumber        decimal.Decimal `xml:"SequenceNumber"`
	AdditionalInformation string          `xml:"AdditionalInformation"`
}

// swagger:model CarbonPriceDue
type CarbonPriceDue struct {
	XMLName             xml.Name        `xml:"CarbonPriceDue"`
	SequenceNumber      decimal.Decimal `xml:"SequenceNumber"`
	InstrumentType      string          `xml:"InstrumentType"`
	LegalActDescription string          `xml:"LegalActDescription"`
	Amount              decimal.Decimal `xml:"Amount"`
	Currency            string          `xml:"Currency"`
	ExchangeRate        string          `xml:"ExchangeRate"`
	EURO                decimal.Decimal `xml:"EURO"`
	Country             string          `xml:"Country"`
	ProductsCovered     ProductCovered  `xml:"ProductsCovered"`
}

// swagger:model ProductCovered
type ProductCovered struct {
	SequenceNumber           decimal.Decimal       `xml:"SequenceNumber"`
	Type                     string                `xml:"Type"`
	CN                       string                `xml:"CN"`
	QuantityCovered          decimal.Decimal       `xml:"QuantityCovered"`
	QuantityCoveredFreeAloc  decimal.Decimal       `xml:"QuantityCoveredFreeAloc"`
	SupplementaryInformation string                `xml:"SupplementaryInformation"`
	AdditionalInformation    string                `xml:"AdditionalInformation"`
	Measure                  ProductCoveredMeasure `xml:"Measure"`
}

// swagger:model ProductCoveredMeasure
type ProductCoveredMeasure struct {
	NetMass         decimal.Decimal `xml:"NetMass"`
	MeasurementUnit string          `xml:"MeasurementUnit"`
}
