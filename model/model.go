package model

import (
	"encoding/xml"
)

//

type QReport struct {
	XMLName               xml.Name                  `xml:"http://xmlns.ec.eu/BusinessObjects/CBAM/Types/V1 QReport"`
	SubmissionDate        DateTimeFull              `xml:"SubmissionDate"`
	ReportId              ConstrainedString         `xml:"ReportId" MinOccurs:"0"`
	ReportingPeriod       ConstrainedString         `xml:"ReportingPeriod" min:"2" max:"2"`
	Year                  ConstrainedInt            `xml:"Year" min:"4" max:"4"`
	TotalImported         ConstrainedInt            `xml:"TotalImported" min:"1" max:"999"`
	Declarant             DeclarantType             `xml:"Declarant"`
	Representative        *RepresentativeType       `xml:"Representative" MinOccurs:"0"`
	Importer              *ImporterType             `xml:"Importer" MinOccurs:"0"`
	NationalCompetentAuth NationalCompetentAuthType `xml:"NationalCompetentAuth"`
	Signatures            SignaturesType            `xml:"Signatures"`
	Remarks               *RemarksType              `xml:"Remarks" MinOccurs:"0"`
	ImportedGood          ImportedGood              `xml:"ImportedGood" MaxOccurs:"99999"`
	ErrorMessages         []string
}

type DeclarantType struct {
	IdentificationNumber ConstrainedString `xml:"IdentificationNumber" min:"1" max:"17"`
	Name                 ConstrainedString `xml:"Name" min:"1" max:"70"`
	Role                 ConstrainedString `xml:"Role" min:"1" max:"5"`
	ActorAddress         AddressType       `xml:"ActorAddress"`
}
type AddressType struct {
	Country              ConstrainedString  `xml:"Country" min:"2" max:"2"`
	SubDivision          *ConstrainedString `xml:"SubDivision" min:"1" max:"70" MinOccurs:"0"`
	City                 ConstrainedString  `xml:"City" min:"1" max:"35"`
	Street               *ConstrainedString `xml:"Street" min:"1" max:"70" MinOccurs:"0"`
	StreetAdditionalLine *ConstrainedString `xml:"StreetAdditionalLine" min:"1" max:"70" MinOccurs:"0"`
	Number               *ConstrainedString `xml:"Number" min:"1" max:"35" MinOccurs:"0"`
	Postcode             *ConstrainedString `xml:"Postcode" min:"1" max:"17" MinOccurs:"0"`
	POBox                *ConstrainedString `xml:"POBox" min:"1" max:"70" MinOccurs:"0"`
}

type ImporterType struct {
	IdentificationNumber ConstrainedString `xml:"IdentificationNumber" min:"1" max:"17"`
	Name                 ConstrainedString `xml:"Name" min:"1" max:"70"`
	ImporterAddress      AddressType       `xml:"ImporterAddress"`
}

type RepresentativeType struct {
	IdentificationNumber  ConstrainedString `xml:"IdentificationNumber" min:"1" max:"17"`
	Name                  ConstrainedString `xml:"Name" min:"1" max:"70"`
	RepresentativeAddress AddressType       `xml:"RepresentativeAddress"`
}

type NationalCompetentAuthType struct {
	ReferenceNumber ConstrainedString `xml:"ReferenceNumber" min:"1" max:"128"`
}

type ReportConfirmationType struct {
	GlobalDataConfirmation  bool              `xml:"GlobalDataConfirmation"`
	UseOfDataConfirmation   bool              `xml:"UseOfDataConfirmation"`
	SignatureDate           SimpleDate        `xml:"SignatureDate"`
	SignaturePlace          ConstrainedString `xml:"SignaturePlace" min:"1" max:"128"`
	Signature               ConstrainedString `xml:"Signature" min:"1" max:"128"`
	PositionOfPersonSending ConstrainedString `xml:"PositionOfPersonSending" min:"1" max:"128"`
}

type SignaturesType struct {
	ReportConfirmation                ReportConfirmationType                 `xml:"ReportConfirmation"`
	ApplicableMethodologyConfirmation *ApplicableMethodologyConfirmationType `xml:"ApplicableMethodologyConfirmation" MinOccurs:"0"`
}

type ApplicableMethodologyConfirmationType struct {
	MethodologyUsed              bool              `xml:"MethodologyUsed"`
	MethodologyConfirmationDate  ConstrainedString `xml:"MethodologyConfirmationDate" min:"10" max:"10"` // Assuming ISO date format (YYYY-MM-DD) for the date.
	MethodologyConfirmationPlace ConstrainedString `xml:"MethodologyConfirmationPlace" min:"1" max:"128"`
}

type RemarksType struct {
	AdditionalInformation ConstrainedString `xml:"AdditionalInformation" min:"1" max:"128"`
}

type ImportedGood struct {
	ItemNumber          ConstrainedInt       `xml:"ItemNumber" min:"1" max:"99"`
	Representative      *RepresentativeType  `xml:"Representative" MinOccurs:"0"` // Optional, no min/max applicable
	Importer            *ImporterType        `xml:"ImporterType" MinOccurs:"0"`   // Optional, no min/max applicable
	CommodityCode       *CommodityCodeType   `xml:"CommodityCode"`                // Optional, no min/max applicable
	OriginCountry       OriginCountryType    `xml:"OriginCountry"`
	ImportedQuantity    ImportedQuantityType `xml:"ImportedQuantity" MinOccurs:"1" MaxOccurs:"9"`
	MeasureImported     MeasureType          `xml:"MeasureImported"`
	TotalEmissions      TotalEmissionsType   `xml:"TotalEmissions"`
	SupportingDocuments SupportingDocument   `xml:"SupportingDocuments" MinOccurs:"0" MaxOccurs:"99"`
	Remarks             *RemarksType         `xml:"Remarks"`
	GoodsEmissions      []GoodsEmissionsType `xml:"GoodsEmissionsType" MinOccurs:"1" MaxOccurs:"999"`
}

type CommodityCodeType struct {
	HsCode           ConstrainedString    `xml:"HsCode" min:"6" max:"6"`
	CnCode           ConstrainedString    `xml:"CnCode" min:"2" max:"2"`
	CommodityDetails CommodityDetailsType `xml:"CommodityDetails"`
}

type CommodityDetailsType struct {
	Description ConstrainedString `xml:"Description" min:"1" max:"512"`
}

type OriginCountryType struct {
	Country ConstrainedString `xml:"Country" min:"2" max:"2"`
}

type ImportedQuantityType struct {
	SequenceNumber           ConstrainedInt         `xml:"SequenceNumber" min:"1" max:"99"`
	Procedure                *ProcedureType         `xml:"Procedure"`
	ImportArea               ImportAreaType         `xml:"ImportArea"`
	MeasureProcedureImported MeasureProcedureType   `xml:"MeasureProcedureType" MaxOccurs:"2"`
	SpecialReferences        *SpecialReferencesType `xml:"SpecialReferences" MinOccurs:"0"`
}

type ProcedureType struct {
	RequestedProc        ConstrainedString        `xml:"RequestedProc" min:"2" max:"2"`
	PreviousProc         *ConstrainedString       `xml:"PreviousProc" min:"2" max:"2" MinOccurs:"0"`
	InwardProcessingInfo InwardProcessingInfoType `xml:"InwardProcessingInfo" MinOccurs:"0" MaxOccurs:"999"`
}

type InwardProcessingInfoType struct {
	MemberStateAuth     ConstrainedString `xml:"MemberStateAuth" min:"2" max:"2"`
	DischargeBillWaiver ConstrainedString `xml:"DischargeBillWaiver" min:"1" max:"1"`
	Authorisation       ConstrainedString `xml:"Authorisation" min:"1" max:"128"`
	StartTime           ConstrainedString `xml:"StartTime" min:"8" max:"8"`
	EndTime             ConstrainedString `xml:"EndTime" min:"8" max:"8"`
	Deadline            ConstrainedString `xml:"Deadline" min:"8" max:"8"`
}

type ImportAreaType struct {
	ImportArea ConstrainedString `xml:"ImportArea" min:"2" max:"2"`
}

type MeasureType struct {
	NetMass            *ConstrainedDecimal `xml:"NetMass"  min:"1" max:"999"`
	SupplementaryUnits *ConstrainedDecimal `xml:"SupplementaryUnits" MinOccurs:"0"`
	MeasurementUnit    ConstrainedString   `xml:"MeasurementUnit" min:"1" max:"5"`
}

type MeasureProcedureType struct {
	NetMass            *ConstrainedDecimal `xml:"NetMass" min:"1" max:"999" MinOccurs:"0"`
	SupplementaryUnits *ConstrainedDecimal `xml:"SupplementaryUnits" MinOccurs:"0"`
	MeasurementUnit    ConstrainedString   `xml:"MeasurementUnit" min:"1" max:"5"`
}

type SpecialReferencesType struct {
	AdditionalInformation ConstrainedString `xml:"AdditionalInformation" min:"1" max:"128"`
}

type TotalEmissionsType struct {
	EmissionsPerUnit ConstrainedDecimal `xml:"EmissionsPerUnit"`
	OverallEmissions ConstrainedDecimal `xml:"OverallEmissions"`
	TotalDirect      ConstrainedDecimal `xml:"TotalDirect"`
	TotalIndirect    ConstrainedDecimal `xml:"TotalIndirect"`
	MeasurementUnit  ConstrainedString  `xml:"MeasurementUnit" min:"5" max:"5"`
}

type SupportingDocument struct {
	SequenceNumber    ConstrainedInt     `xml:"SequenceNumber" min:"1" max:"99"`
	Type              ConstrainedString  `xml:"Type" min:"1" max:"8"`
	Country           *ConstrainedString `xml:"Country" min:"2" max:"2" MinOccurs:"0"` // Optional
	ReferenceNumber   ConstrainedString  `xml:"ReferenceNumber" min:"1" max:"70"`
	LineItemNumber    *ConstrainedString `xml:"LineItemNumber" min:"5" max:"5" MinOccurs:"0"`   // Assuming NUMERIC5 implies length of 5
	IssuingAuthName   *ConstrainedString `xml:"IssuingAuthName" min:"1" max:"70" MinOccurs:"0"` // Optional
	ValidityStartDate SimpleDate         `xml:"ValidityStartDate" MinOccurs:"0"`                // Optional
	ValidityEndDate   SimpleDate         `xml:"ValidityEndDate" MinOccurs:"0"`                  // Optional
	Description       *ConstrainedString `xml:"Description" min:"1" max:"128" MinOccurs:"0"`    // Optional
	Attachment        *AttachmentType    `xml:"Attachment" MinOccurs:"0"`                       // Assuming AttachmentType is defined elsewhere.
}

type AttachmentType struct {
	Filename ConstrainedString  `min:"1" max:"128"`
	URI      *ConstrainedString `min:"1" max:"128" MinOccurs:"0"`
	MIME     ConstrainedString  `min:"1" max:"700"`
	Binary   byte               `json:"binary,omitempty"`
}
type GoodsEmissionsType struct {
	SequenceNumber             ConstrainedInt                 `xml:"SequenceNumber" min:"1" max:"99"`
	ProductionCountry          *ConstrainedString             `xml:"ProductionCountry" min:"2" max:"2" MinOccurs:"0"` // Optional
	InstallationOperator       *InstallationOperatorType      `xml:"InstallationOperator" MinOccurs:"0"`              // Optional
	Installation               *InstallationType              `xml:"Installation" MinOccurs:"0"`                      // Optional
	ProducedMeasure            MeasureType                    `xml:"ProducedMeasure"`
	InstallationEmissions      InstallationEmissionsType      `xml:"InstallationEmissions"`
	DirectEmissions            DirectEmissionsType            `xml:"DirectEmissions"`
	IndirectEmissions          *IndirectEmissionsType         `xml:"IndirectEmissions" MinOccurs:"0"` // Optional
	ProdMethodQualifyingParams ProdMethodQualifyingParamsType `xml:"ProdMethodQualifyingParams" MaxOccurs:"99"`
	SupportingDocuments        SupportingDocument             `xml:"SupportingDocuments" MaxOccurs:"99"`
	CarbonPriceDue             CarbonPriceDueType             `xml:"CarbonPriceDue" MaxOccurs:"99"`
	RemarksEmissions           RemarksEmissionsType           `xml:"RemarksEmissions" MaxOccurs:"9"`
}

type InstallationOperatorType struct {
	OperatorId      ConstrainedString  `xml:"OperatorId" min:"1" max:"17"`
	OperatorName    ConstrainedString  `xml:"OperatorName" min:"1" max:"70"`
	OperatorAddress AddressType        `xml:"OperatorAddress"`
	ContactDetails  ContactDetailsType `xml:"ContactDetails" MaxOccurs:"9"`
}

type InstallationType struct {
	InstallationId   ConstrainedString       `xml:"InstallationId" min:"1" max:"17"`
	InstallationName ConstrainedString       `xml:"InstallationName" min:"1" max:"128"`
	EconomicActivity *ConstrainedString      `xml:"EconomicActivity" min:"1" max:"128" MinOccurs:"0"` // Optional
	Address          InstallationAddressType `xml:"Address"`
}

type InstallationAddressType struct {
	EstablishmentCountry ConstrainedString  `xml:"EstablishmentCountry" min:"2" max:"2"`
	SubDivision          *ConstrainedString `xml:"SubDivision" min:"1" max:"35" MinOccurs:"0"`          // Optional
	City                 *ConstrainedString `xml:"City" min:"1" max:"35" MinOccurs:"0"`                 // Optional
	Street               *ConstrainedString `xml:"Street" min:"1" max:"70" MinOccurs:"0"`               // Optional
	StreetAdditionalLine *ConstrainedString `xml:"StreetAdditionalLine" min:"1" max:"70" MinOccurs:"0"` // Optional
	Number               *ConstrainedString `xml:"Number" min:"1" max:"35" MinOccurs:"0"`               // Optional
	Postcode             *ConstrainedString `xml:"Postcode" min:"1" max:"17" MinOccurs:"0"`             // Optional
	POBox                *ConstrainedString `xml:"POBox" min:"1" max:"70" MinOccurs:"0"`                // Optional
	PlotParcelNumber     *ConstrainedString `xml:"PlotParcelNumber" min:"1" max:"15" MinOccurs:"0"`     // Optional
	UNLOCODE             *ConstrainedString `xml:"UNLOCODE" min:"1" max:"17" MinOccurs:"0"`             // Optional
	Latitude             *ConstrainedString `xml:"Latitude" min:"1" max:"17" MinOccurs:"0"`             // Optional
	Longitude            *ConstrainedString `xml:"Longitude" min:"1" max:"17" MinOccurs:"0"`            // Optional
	CoordinatesType      *ConstrainedString `xml:"CoordinatesType" min:"1" max:"5" MinOccurs:"0"`       // Optional
}

type DirectEmissionsType struct {
	DeterminationType                  *ConstrainedString  `xml:"DeterminationType" min:"1" max:"5"`            // Optional
	DeterminationTypeElectricity       *ConstrainedString  `xml:"DeterminationTypeElectricity" min:"1" max:"5"` // Optional
	ApplicableReportingTypeMethodology ConstrainedString   `xml:"ApplicableReportingTypeMethodology" min:"1" max:"5"`
	ApplicableReportingMethodology     *ConstrainedString  `xml:"ApplicableReportingMethodology" min:"1" max:"256"`   // Optional
	SpecificEmbeddedEmissions          *ConstrainedString  `xml:"SpecificEmbeddedEmissions" min:"1" max:"165"`        // Optional, represent DECIMAL165 as string
	OtherSourceIndication              *ConstrainedString  `xml:"OtherSourceIndication" min:"1" max:"256"`            // Optional
	EmissionFactorSourceElectricity    *ConstrainedString  `xml:"EmissionFactorSourceElectricity" min:"1" max:"5"`    // Optional
	EmissionFactor                     *ConstrainedString  `xml:"EmissionFactor" min:"1" max:"165"`                   // Optional, represent DECIMAL165 as string
	ElectricityImported                *ConstrainedDecimal `xml:"ElectricityImported"`                                // Optional
	TotalEmbeddedElectricityImported   *ConstrainedString  `xml:"TotalEmbeddedElectricityImported" min:"1" max:"165"` // Optional, represent DECIMAL165 as string
	MeasurementUnit                    ConstrainedString   `xml:"MeasurementUnit" min:"1" max:"5"`
	EmissionFactorSourceValue          *ConstrainedString  `xml:"EmissionFactorSourceValue" min:"1" max:"512"` // Optional
	Justification                      *ConstrainedString  `xml:"Justification" min:"1" max:"512"`             // Optional
	ConditionalityFulfillment          *ConstrainedString  `xml:"ConditionalityFulfillment" min:"1" max:"512"` // Optional
}

type ContactDetailsType struct {
	Name  ConstrainedString `xml:"Name" min:"1" max:"70"`
	Phone ConstrainedString `xml:"Phone" min:"1" max:"35"`
	Email ConstrainedString `xml:"Email" min:"1" max:"256"`
}

type ProdMethodQualifyingParamsType struct {
	SequenceNumber               ConstrainedInt           `xml:"SequenceNumber" min:"1" max:"99"`
	MethodId                     ConstrainedString        `xml:"MethodId" min:"1" max:"5"`
	MethodName                   ConstrainedString        `xml:"MethodName" min:"1" max:"256"`
	SteelMillNumber              *ConstrainedString       `xml:"SteelMillNumber" min:"1" max:"256" MinOccurs:"0"`       // Optional
	AdditionalInformation        *ConstrainedString       `xml:"AdditionalInformation" min:"1" max:"512" MinOccurs:"0"` // Optional
	DirectQualifyingParameters   QualifyingParametersType `xml:"DirectQualifyingParameters" MinOccurs:"0" MaxOccurs:"99"`
	IndirectQualifyingParameters QualifyingParametersType `xml:"IndirectQualifyingParameters" MinOccurs:"0" MaxOccurs:"99"`
}

type QualifyingParametersType struct {
	SequenceNumber        ConstrainedInt     `xml:"SequenceNumber" min:"1" max:"99"`
	DeterminationType     ConstrainedString  `xml:"DeterminationType" min:"1" max:"5"`
	ParameterId           ConstrainedString  `xml:"ParameterId" min:"1" max:"5"`
	ParameterName         ConstrainedString  `xml:"ParameterName" min:"1" max:"256"`
	Description           *ConstrainedString `xml:"Description" min:"1" max:"256" MinOccurs:"0"` // Optional
	ParameterValueType    ConstrainedString  `xml:"ParameterValueType" min:"1" max:"256"`
	ParameterValue        ConstrainedString  `xml:"ParameterValue" min:"1" max:"256"`
	AdditionalInformation *ConstrainedString `xml:"AdditionalInformation" min:"1" max:"512" MinOccurs:"0"` // Optional
}

type InstallationEmissionsType struct {
	OverallEmissions ConstrainedDecimal `xml:"OverallEmissions"`
	TotalDirect      ConstrainedDecimal `xml:"TotalDirect"`
	TotalIndirect    ConstrainedDecimal `xml:"TotalIndirect"`
	MeasurementUnit  ConstrainedString  `xml:"MeasurementUnit" min:"1" max:"5"`
}

type IndirectEmissionsType struct {
	DeterminationType         ConstrainedString  `xml:"DeterminationType" min:"1" max:"5"`
	EmissionFactorSource      *ConstrainedString `xml:"EmissionFactorSource" min:"1" max:"5"` // Optional
	EmissionFactor            *ConstrainedString `xml:"EmissionFactor" min:"1" max:"165"`     // Optional, DECIMAL165 represented as string
	SpecificEmbeddedEmissions ConstrainedDecimal `xml:"SpecificEmbeddedEmissions"`
	MeasurementUnit           ConstrainedString  `xml:"MeasurementUnit" min:"1" max:"5"`
	ElectricityConsumed       *ConstrainedString `xml:"ElectricityConsumed" min:"1" max:"52"` // Optional, DECIMAL52 represented as string
	ElectricitySource         ConstrainedString  `xml:"ElectricitySource" min:"1" max:"5"`
	OtherSourceIndication     *ConstrainedString `xml:"OtherSourceIndication" min:"1" max:"256"`     // Optional
	EmissionFactorSourceValue *ConstrainedString `xml:"EmissionFactorSourceValue" min:"1" max:"512"` // Optional
}

type CarbonPriceDueType struct {
	SequenceNumber      ConstrainedInt      `xml:"SequenceNumber" min:"1" max:"99"`
	InstrumentType      ConstrainedString   `xml:"InstrumentType" min:"1" max:"5"`
	LegalActDescription ConstrainedString   `xml:"LegalActDescription" min:"1" max:"512"`
	Amount              ConstrainedDecimal  `xml:"Amount"`
	Currency            ConstrainedString   `xml:"Currency" min:"1" max:"3"`
	ExchangeRate        ConstrainedString   `xml:"ExchangeRate" min:"1" max:"3"`
	EURO                ConstrainedDecimal  `xml:"EURO"`
	Country             ConstrainedString   `xml:"Country" min:"2" max:"2"`
	ProductsCovered     ProductsCoveredType `xml:"ProductsCovered"`
}

type ProductsCoveredType struct {
	SequenceNumber           ConstrainedInt     `xml:"SequenceNumber" min:"1" max:"99"`
	Type                     ConstrainedString  `xml:"Type" min:"1" max:"8"`
	CN                       *ConstrainedString `xml:"CN" min:"1" max:"8"`                         // Optional, assuming NUMERIC8 can be represented as string
	QuantityCovered          ConstrainedString  `xml:"QuantityCovered" min:"1" max:"165"`          // DECIMAL165 as string
	QuantityCoveredFreeAloc  ConstrainedString  `xml:"QuantityCoveredFreeAloc" min:"1" max:"165"`  // DECIMAL165 as string
	SupplementaryInformation *ConstrainedString `xml:"SupplementaryInformation" min:"1" max:"256"` // Optional
	AdditionalInformation    *ConstrainedString `xml:"AdditionalInformation" min:"1" max:"512"`    // Optional
	Measure                  MeasureType        `xml:"Measure"`
}

type RemarksEmissionsType struct {
	SequenceNumber        ConstrainedInt    `xml:"SequenceNumber" min:"1" max:"99"`
	AdditionalInformation ConstrainedString `xml:"AdditionalInformation" min:"1" max:"512"`
}

type Error struct {
	StatusCode  int         `json:"-"`
	Message     string      `json:"message"`
	ErrorDetail string      `json:"error,omitempty"`
	Hints       []string    `json:"hints,omitempty"`
	Details     interface{} `json:"details,omitempty"`
}
