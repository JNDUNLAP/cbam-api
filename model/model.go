package model

import "time"

type DateTime time.Time
type XSDecimal float64
type XSBoolean bool

type DECIMAL52 float64
type DECIMAL165 float64
type DECIMAL166 float64

type STRING3 string
type STRING2 string
type STRING4 string
type STRING6 string
type STRING5 string
type STRING8 string
type STRING15 string
type STRING17 string
type STRING22 string
type STRING35 string
type STRING70 string
type STRING128 string
type STRING256 string
type STRING512 string

type NUMERIC1 int
type NUMERIC2 int
type NUMERIC4 int
type NUMERIC5 int
type NUMERIC8 int

type SimpleDate struct {
	time.Time
}

func (d SimpleDate) FormatISO8601() string {
	return d.Format("2006-01-02T15:04:05Z")
}

type Error struct {
	StatusCode  int         `json:"-"`
	Message     string      `json:"message"`
	ErrorDetail string      `json:"error,omitempty"`
	Hints       []string    `json:"hints,omitempty"`
	Details     interface{} `json:"details,omitempty"`
}

type QReport struct {
	// XMLName               xml.Name `xml:"QReport"`
	SubmissionDate DateTime
	DraftReportId  STRING22
	// ReportId              STRING22
	ReportingPeriod       STRING2
	Year                  NUMERIC4
	TotalImported         int
	TotalEmissions        DECIMAL165
	Declarant             DeclarantType
	Representative        *RepresentativeType
	Importer              *ImporterType
	NationalCompetentAuth NationalCompetentAuthType
	Signatures            SignaturesType
	Remarks               *RemarksType
	ImportedGood          ImportedGood
}

type DeclarantType struct {
	IdentificationNumber STRING17
	Name                 STRING70
	Role                 STRING5
	ActorAddress         AddressType
}

type ImporterType struct {
	IdentificationNumber STRING17
	Name                 STRING70
	ImporterAddress      AddressType
}

type RepresentativeType struct {
	IdentificationNumber  STRING17
	Name                  STRING70
	RepresentativeAddress AddressType
}
type AddressType struct {
	Country              STRING2
	SubDivision          *STRING35 // Optional field
	City                 STRING35
	Street               *STRING70 // Optional field
	StreetAdditionalLine *STRING70 // Optional field
	Number               *STRING35 // Optional field
	Postcode             *STRING17 // Optional field
	POBox                *STRING70 // Optional field
}

type NationalCompetentAuthType struct {
	ReferenceNumber STRING128
}

type SignaturesType struct {
	ReportConfirmation                ReportConfirmationType
	ApplicableMethodologyConfirmation *ApplicableMethodologyConfirmationType
}

type ReportConfirmationType struct {
	GlobalDataConfirmation  XSBoolean
	UseOfDataConfirmation   XSBoolean
	SignatureDate           NUMERIC8
	SignaturePlace          STRING128
	Signature               STRING128
	PositionOfPersonSending STRING128
}

type ApplicableMethodologyConfirmationType struct {
	OtherApplicableReportingMethodology XSBoolean
}

type RemarksType struct {
	AdditionalInformation STRING128
}

type ImportedGood struct {
	ItemNumber          int
	Representative      *RepresentativeType
	Importer            *ImporterType
	CommodityCode       *CommodityCodeType
	OriginCountry       OriginCountryType
	ImportedQuantity    ImportedQuantityType
	MeasureImported     MeasureType
	TotalEmissions      TotalEmissionsType
	SupportingDocuments SupportingDocument
	Remarks             *RemarksType
	GoodsEmissions      GoodsEmissionsType
}
type CommodityCodeType struct {
	HsCode           STRING6
	CnCode           STRING2
	CommodityDetails CommodityDetailsType
}

type CommodityDetailsType struct {
	Description STRING512
}

type OriginCountryType struct {
	Country STRING2
}

type ImportedQuantityType struct {
	SequenceNumber           int
	Procedure                *ProcedureType
	ImportArea               ImportAreaType
	MeasureProcedureImported MeasureProcedureType
	SpecialReferences        *SpecialReferencesType
}

type ProcedureType struct {
	RequestedProc        STRING2
	PreviousProc         *STRING2
	InwardProcessingInfo InwardProcessingInfoType
}

type InwardProcessingInfoType struct {
	MemberStateAuth     STRING2
	DischargeBillWaiver NUMERIC1
	Authorisation       STRING128
	StartTime           NUMERIC8
	EndTime             NUMERIC8
	Deadline            NUMERIC8
}

type ImportAreaType struct {
	ImportArea STRING5
}

type MeasureType struct {
	NetMass            *DECIMAL166
	SupplementaryUnits *DECIMAL166
	MeasurementUnit    STRING5
}

type MeasureProcedureType struct {
	Indicator          NUMERIC1
	NetMass            *DECIMAL166
	SupplementaryUnits *DECIMAL166
	MeasurementUnit    STRING5
}

type SpecialReferencesType struct {
	AdditionalInformation STRING128
}

type TotalEmissionsType struct {
	EmissionsPerUnit XSDecimal
	OverallEmissions XSDecimal
	TotalDirect      XSDecimal
	TotalIndirect    XSDecimal
	MeasurementUnit  STRING5
}

type SupportingDocument struct {
	SequenceNumber    int
	Type              STRING8
	Country           *STRING2
	ReferenceNumber   STRING70
	LineItemNumber    *NUMERIC5
	IssuingAuthName   *STRING70
	ValidityStartDate *DateTime
	ValidityEndDate   *DateTime
	Description       *STRING128
	Attachment        *AttachmentType
}

type AttachmentType struct {
	Filename STRING128
	URI      *STRING128
	MIME     STRING70
	Binary   byte // base64Binary maps to byte in Go
}

type GoodsEmissionsType struct {
	SequenceNumber             int
	ProductionCountry          *STRING2
	InstallationOperator       *InstallationOperatorType
	Installation               *InstallationType
	ProducedMeasure            MeasureType
	InstallationEmissions      InstallationEmissionsType
	DirectEmissions            DirectEmissionsType
	IndirectEmissions          *IndirectEmissionsType
	ProdMethodQualifyingParams ProdMethodQualifyingParamsType
	SupportingDocuments        SupportingDocument
	CarbonPriceDue             CarbonPriceDueType
	RemarksEmissions           RemarksEmissionsType
}

type InstallationOperatorType struct {
	OperatorId      STRING17
	OperatorName    STRING70
	OperatorAddress AddressType
	ContactDetails  ContactDetailsType
}

type InstallationType struct {
	InstallationId   STRING17
	InstallationName STRING128
	EconomicActivity *STRING128
	Address          InstallationAddressType
}

type InstallationAddressType struct {
	EstablishmentCountry STRING2
	SubDivision          *STRING35
	City                 *STRING35
	Street               *STRING70
	StreetAdditionalLine *STRING70
	Number               *STRING35
	Postcode             *STRING17
	POBox                *STRING70
	PlotParcelNumber     *STRING15
	UNLOCODE             *STRING17
	Latitude             *STRING17
	Longitude            *STRING17
	CoordinatesType      *STRING5
}

type ContactDetailsType struct {
	Name  STRING70
	Phone STRING35
	Email STRING256
}

type ProdMethodQualifyingParamsType struct {
	SequenceNumber               int
	MethodId                     STRING5
	MethodName                   STRING256
	SteelMillNumber              *STRING256
	AdditionalInformation        *STRING512
	DirectQualifyingParameters   QualifyingParametersType
	IndirectQualifyingParameters QualifyingParametersType
}

type QualifyingParametersType struct {
	SequenceNumber        int
	DeterminationType     STRING5
	ParameterId           STRING5
	ParameterName         STRING256
	Description           *STRING256
	ParameterValueType    STRING256
	ParameterValue        STRING256
	AdditionalInformation *STRING512
}

type InstallationEmissionsType struct {
	OverallEmissions XSDecimal
	TotalDirect      XSDecimal
	TotalIndirect    XSDecimal
	MeasurementUnit  STRING5
}

type DirectEmissionsType struct {
	DeterminationType                  *STRING5
	DeterminationTypeElectricity       *STRING5
	ApplicableReportingTypeMethodology STRING5
	ApplicableReportingMethodology     *STRING256
	SpecificEmbeddedEmissions          *DECIMAL165
	OtherSourceIndication              *STRING256
	EmissionFactorSourceElectricity    *STRING5
	EmissionFactor                     *DECIMAL165
	ElectricityImported                *XSDecimal
	TotalEmbeddedElectricityImported   *DECIMAL165
	MeasurementUnit                    STRING5
	EmissionFactorSourceValue          *STRING512
	Justification                      *STRING512
	ConditionalityFulfillment          *STRING512
}

type IndirectEmissionsType struct {
	DeterminationType         STRING5
	EmissionFactorSource      *STRING5
	EmissionFactor            *DECIMAL165
	SpecificEmbeddedEmissions XSDecimal
	MeasurementUnit           STRING5
	ElectricityConsumed       *DECIMAL52
	ElectricitySource         STRING5
	OtherSourceIndication     *STRING256
	EmissionFactorSourceValue *STRING512
}

type CarbonPriceDueType struct {
	SequenceNumber      int
	InstrumentType      STRING5
	LegalActDescription STRING512
	Amount              XSDecimal
	Currency            STRING3
	ExchangeRate        STRING3
	EURO                XSDecimal
	Country             STRING2
	ProductsCovered     ProductsCoveredType
}

type ProductsCoveredType struct {
	SequenceNumber           int
	Type                     STRING8
	CN                       *NUMERIC8
	QuantityCovered          DECIMAL165
	QuantityCoveredFreeAloc  DECIMAL165
	SupplementaryInformation *STRING256
	AdditionalInformation    *STRING512
	Measure                  MeasureType
}

type RemarksEmissionsType struct {
	SequenceNumber        int
	AdditionalInformation STRING512
}
