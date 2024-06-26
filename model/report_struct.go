package model

import "time"

type Report struct {
	Submissiondate                     time.Time `json:"SubmissionDate" validate:"regex="[0-9]{4}-[0-9][0-9]-[0-9][0-9]T[0-9][0-9]:[0-9][0-9]:[0-9][0-9]Z""`
	Draftreportid                      string    `json:"DraftReportId" validate:"min=1,max=22"`
	Reportid                           string    `json:"ReportId" validate:"min=1,max=22"`
	Reportingperiod                    string    `json:"ReportingPeriod" validate:"min=2,max=2"`
	Year                               string    `json:"Year" validate:"max=9999"`
	Totalimported                      string    `json:"TotalImported" validate:"regex="[0-9]{1,5}""`
	Totalemissions                     string    `json:"TotalEmissions" validate:"max=9999999999999999"`
	Identificationnumber               string    `json:"IdentificationNumber" validate:"min=1,max=17"`
	Name                               string    `json:"Name" validate:"min=1,max=70"`
	Role                               string    `json:"Role" validate:"min=1,max=5"`
	Referencenumber                    string    `json:"ReferenceNumber" validate:"min=1,max=70"`
	Signaturedate                      string    `json:"SignatureDate" validate:"max=99999999"`
	Signatureplace                     string    `json:"SignaturePlace" validate:"min=1,max=128"`
	Signature                          string    `json:"Signature" validate:"min=1,max=128"`
	Positionofpersonsending            string    `json:"PositionOfPersonSending" validate:"min=1,max=128"`
	Additionalinformation              string    `json:"AdditionalInformation" validate:"min=1,max=512"`
	Itemnumber                         string    `json:"ItemNumber" validate:"regex="[0-9]{1,5}""`
	Hscode                             string    `json:"HsCode" validate:"min=6,max=6"`
	Cncode                             string    `json:"CnCode" validate:"min=2,max=2"`
	Description                        string    `json:"Description" validate:"min=1,max=256"`
	Country                            string    `json:"Country" validate:"min=2,max=2"`
	Sequencenumber                     string    `json:"SequenceNumber" validate:"regex="[0-9]{1,5}""`
	Requestedproc                      string    `json:"RequestedProc" validate:"min=2,max=2"`
	Previousproc                       string    `json:"PreviousProc" validate:"min=2,max=2"`
	Memberstateauth                    string    `json:"MemberStateAuth" validate:"min=2,max=2"`
	Dischargebillwaiver                string    `json:"DischargeBillWaiver" validate:"max=9"`
	Authorisation                      string    `json:"Authorisation" validate:"min=1,max=512"`
	Starttime                          string    `json:"StartTime" validate:"max=99999999"`
	Endtime                            string    `json:"EndTime" validate:"max=99999999"`
	Deadline                           string    `json:"Deadline" validate:"max=99999999"`
	Importarea                         string    `json:"ImportArea" validate:"min=1,max=5"`
	Netmass                            string    `json:"NetMass" validate:"max=9999999999999999"`
	Supplementaryunits                 string    `json:"SupplementaryUnits" validate:"max=9999999999999999"`
	Measurementunit                    string    `json:"MeasurementUnit" validate:"min=1,max=5"`
	Indicator                          string    `json:"Indicator" validate:"max=9"`
	Type                               string    `json:"Type" validate:"min=1,max=5"`
	Lineitemnumber                     string    `json:"LineItemNumber" validate:"regex="[0-9]{1,5}""`
	Issuingauthname                    string    `json:"IssuingAuthName" validate:"min=1,max=70"`
	Validitystartdate                  string    `json:"ValidityStartDate" validate:"regex="[0-9]{4}-[0-9][0-9]-[0-9][0-9]""`
	Validityenddate                    string    `json:"ValidityEndDate" validate:"regex="[0-9]{4}-[0-9][0-9]-[0-9][0-9]""`
	Filename                           string    `json:"Filename" validate:"min=1,max=256"`
	Uri                                string    `json:"URI" validate:"min=1,max=2048"`
	Mime                               string    `json:"MIME" validate:"min=1,max=71"`
	Determinationtype                  string    `json:"DeterminationType" validate:"min=1,max=5"`
	Parameterid                        string    `json:"ParameterId" validate:"min=1,max=5"`
	Parametername                      string    `json:"ParameterName" validate:"min=1,max=256"`
	Parametervaluetype                 string    `json:"ParameterValueType" validate:"min=1,max=256"`
	Parametervalue                     string    `json:"ParameterValue" validate:"min=1,max=256"`
	Productioncountry                  string    `json:"ProductionCountry" validate:"min=2,max=2"`
	Installationid                     string    `json:"InstallationId" validate:"min=1,max=17"`
	Installationname                   string    `json:"InstallationName" validate:"min=1,max=256"`
	Economicactivity                   string    `json:"EconomicActivity" validate:"min=1,max=256"`
	Establishmentcountry               string    `json:"EstablishmentCountry" validate:"min=2,max=2"`
	Subdivision                        string    `json:"SubDivision" validate:"min=1,max=35"`
	City                               string    `json:"City" validate:"min=1,max=35"`
	Street                             string    `json:"Street" validate:"min=1,max=70"`
	Streetadditionalline               string    `json:"StreetAdditionalLine" validate:"min=1,max=70"`
	Number                             string    `json:"Number" validate:"min=1,max=35"`
	Postcode                           string    `json:"Postcode" validate:"min=1,max=17"`
	Pobox                              string    `json:"POBox" validate:"min=1,max=70"`
	Plotparcelnumber                   string    `json:"PlotParcelNumber" validate:"min=1,max=15"`
	Unlocode                           string    `json:"UNLOCODE" validate:"min=1,max=17"`
	Latitude                           string    `json:"Latitude" validate:"min=1,max=17"`
	Longitude                          string    `json:"Longitude" validate:"min=1,max=17"`
	Coordinatestype                    string    `json:"CoordinatesType" validate:"min=1,max=5"`
	Operatorid                         string    `json:"OperatorId" validate:"min=1,max=17"`
	Operatorname                       string    `json:"OperatorName" validate:"min=1,max=70"`
	Phone                              string    `json:"Phone" validate:"min=1,max=35"`
	Email                              string    `json:"Email" validate:"min=1,max=256"`
	Methodid                           string    `json:"MethodId" validate:"min=1,max=5"`
	Methodname                         string    `json:"MethodName" validate:"min=1,max=256"`
	Steelmillnumber                    string    `json:"SteelMillNumber" validate:"min=1,max=256"`
	Determinationtypeelectricity       string    `json:"DeterminationTypeElectricity" validate:"min=1,max=5"`
	Applicablereportingtypemethodology string    `json:"ApplicableReportingTypeMethodology" validate:"min=1,max=5"`
	Applicablereportingmethodology     string    `json:"ApplicableReportingMethodology" validate:"min=1,max=256"`
	Specificembeddedemissions          string    `json:"SpecificEmbeddedEmissions" validate:"max=9999999999999999"`
	Othersourceindication              string    `json:"OtherSourceIndication" validate:"min=1,max=256"`
	Emissionfactorsourceelectricity    string    `json:"EmissionFactorSourceElectricity" validate:"min=1,max=5"`
	Emissionfactor                     string    `json:"EmissionFactor" validate:"max=9999999999999999"`
	Totalembeddedelectricityimported   string    `json:"TotalEmbeddedElectricityImported" validate:"max=9999999999999999"`
	Emissionfactorsourcevalue          string    `json:"EmissionFactorSourceValue" validate:"min=1,max=512"`
	Justification                      string    `json:"Justification" validate:"min=1,max=512"`
	Conditionalityfulfillment          string    `json:"ConditionalityFulfillment" validate:"min=1,max=512"`
	Emissionfactorsource               string    `json:"EmissionFactorSource" validate:"min=1,max=5"`
	Electricityconsumed                string    `json:"ElectricityConsumed" validate:"max=9999999"`
	Electricitysource                  string    `json:"ElectricitySource" validate:"min=1,max=5"`
	Instrumenttype                     string    `json:"InstrumentType" validate:"min=1,max=5"`
	Legalactdescription                string    `json:"LegalActDescription" validate:"min=1,max=512"`
	Currency                           string    `json:"Currency" validate:"min=3,max=3"`
	Exchangerate                       string    `json:"ExchangeRate" validate:"min=1,max=5"`
	Cn                                 string    `json:"CN" validate:"min=1,max=8"`
	Quantitycovered                    string    `json:"QuantityCovered" validate:"max=9999999999999999"`
	Quantitycoveredfreealoc            string    `json:"QuantityCoveredFreeAloc" validate:"max=9999999999999999"`
	Supplementaryinformation           string    `json:"SupplementaryInformation" validate:"min=1,max=256"`
}
