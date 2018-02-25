package Types

import (
	"time"
)

type anySimpleType string

/***********************************************************
   					Non-derived types
 ***********************************************************/


/*
		The string datatype represents character strings in XML.
		The ·value space· of string is the set of finite-length sequences of characters.
		String has the following constraining facets:
		• length
		• minLength
		• maxLength
		• pattern
		• enumeration
		• whiteSpace

		More info: https://www.w3.org/TR/xmlschema-2/#string

		//TODO: valid/invalid character declaration and process restrictions
 */
type String string

/*
	Construct an instance of xsd String type
 */
func (tp String) NewString(data string) String {
	return String(data)
}

/*
		Boolean has the ·value space· required to support the mathematical concept of binary-valued logic: {true, false}.
		Boolean has the following ·constraining facets·:
		• pattern
		• whiteSpace

		More info: https://www.w3.org/TR/xmlschema-2/#boolean

		//TODO: process restrictions
 */
type Boolean bool

/*
	Construct an instance of xsd Boolean type
 */
func (tp Boolean) NewBool(data bool) Boolean {
	return Boolean(data)
}

/*
		Float is patterned after the IEEE single-precision 32-bit floating point type
		Float has the following ·constraining facets·:
		• pattern
		• enumeration
		• whiteSpace
		• maxInclusive
		• maxExclusive
		• minInclusive
		• minExclusive

		More info: https://www.w3.org/TR/xmlschema-2/#float

		//TODO: process restrictions
 */
type Float float32

/*
	Construct an instance of xsd Float type
 */
func (tp Float) NewFloat(data float32) Float {
	return Float(data)
}

/*
		The double datatype is patterned after the IEEE double-precision 64-bit floating point type
		Double has the following ·constraining facets·:
		• pattern
		• enumeration
		• whiteSpace
		• maxInclusive
		• maxExclusive
		• minInclusive
		• minExclusive

		More info: https://www.w3.org/TR/xmlschema-2/#double

		//TODO: process restrictions
 */
type Double float64

/*
	Construct an instance of xsd Double type
 */
func (tp Double) NewDouble(data float64) Double {
	return Double(data)
}

/*
		The type decimal represents a decimal number of arbitrary precision.
		Schema processors vary in the number of significant digits they support,
		but a conforming processor must support a minimum of 18 significant digits.
		The format of xsd:decimal is a sequence of digits optionally preceded by a sign ("+" or "-")
		and optionally containing a period. The value may start or end with a period.
		If the fractional part is 0 then the period and trailing zeros may be omitted.
		Leading and trailing zeros are permitted, but they are not considered significant.
		That is, the decimal values 3.0 and 3.0000 are considered equal.

		Source: http://www.datypic.com/sc/xsd/t-xsd_decimal.html

		Decimal has the following ·constraining facets·:
		• totalDigits
		• fractionDigits
		• pattern
		• whiteSpace
		• enumeration
		• maxInclusive
		• maxExclusive
		• minInclusive
		• minExclusive

		More info: https://www.w3.org/TR/xmlschema-2/#decimal

		//TODO: process restrictions, valid/invalid characters(commas are not permitted; the decimal separator must be a period)

 */
type Decimal string

/*
	Construct an instance of xsd Decimal type
 */
func (tp Decimal) NewDecimal(data string) Decimal {
	return Decimal(data)
}

//duration is the [ISO 8601] extended format PnYn MnDTnH nMnS,
// where nY represents the number of years, nM the number of months,
// nD the number of days, 'T' is the date/time separator,
// nH the number of hours, nM the number of minutes and nS the number of seconds.
// The number of seconds can include decimal digits to arbitrary precision.
//PnYnMnDTnHnMnS
type Duration anySimpleType

func (tp Duration) NewDateTime(t time.Time) Duration {
	return Duration(t.Format("2006-01-02T15:04:05-0700"))
}

//TODO: decide good type for time with proper format
//The ·lexical space· of dateTime consists of finite-length sequences of characters of the form:
// '-'? yyyy '-' mm '-' dd 'T' hh ':' mm ':' ss ('.' s+)? (zzzzzz)?,
type DateTime anySimpleType

//func (tp ) NewDateTime() {
//
//}

//type Time anySimpleType
//
//type Date anySimpleType
//
//type GYearMonth anySimpleType
//
//type GYear anySimpleType
//
//type GMonthDay anySimpleType
//
//type GDay anySimpleType
//
//type GMonth anySimpleType
//
//type HexBinary anySimpleType
//
//type Base64Binary anySimpleType
//
//type AnyURI anySimpleType
//
//type QName anySimpleType
//
//type NOTATION anySimpleType

/*
   Derived types
 */



