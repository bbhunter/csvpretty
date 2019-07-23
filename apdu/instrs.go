package apdu

// Code generated by "go run github.com/zemnmez/cardauth/apdu/gen/instrs instrs.go" DO NOT EDIT.

//go:generate go run github.com/zemnmez/cardauth/apdu/gen/instrs $GOFILE
//go:generate gofmt -w -s $GOFILE

import "fmt"

// Instruction represents a smart card instruction as defined in ISO/IEC 7816-4:2005(E)
type Instruction string

const (
	// InstrDeactivateFile represents the "DEACTIVATE FILE" instruction as defined in ISO/IEC 7816-4:2005(E) Part 9
	InstrDeactivateFile Instruction = "\x04"

	// InstrEraseRecord represents the "ERASE RECORD (S)" instruction as defined in ISO/IEC 7816-4:2005(E) 7.3.8
	InstrEraseRecord Instruction = "\x0C"

	// InstrEraseBinary represents the "ERASE BINARY" instruction as defined in ISO/IEC 7816-4:2005(E) 7.2.7
	InstrEraseBinary Instruction = "\x0E\x0F"

	// InstrPerformScqlOperation represents the "PERFORM SCQL OPERATION" instruction as defined in ISO/IEC 7816-4:2005(E) Part 7
	InstrPerformScqlOperation Instruction = "\x10"

	// InstrPerformTransactionOperation represents the "PERFORM TRANSACTION OPERATION" instruction as defined in ISO/IEC 7816-4:2005(E) Part 7
	InstrPerformTransactionOperation Instruction = "\x12"

	// InstrPerformUserOperation represents the "PERFORM USER OPERATION" instruction as defined in ISO/IEC 7816-4:2005(E) Part 7
	InstrPerformUserOperation Instruction = "\x14"

	// InstrVerify represents the "VERIFY" instruction as defined in ISO/IEC 7816-4:2005(E) 7.5.6
	InstrVerify Instruction = "\x20\x21"

	// InstrManageSecurityEnvironment represents the "MANAGE SECURITY ENVIRONMENT" instruction as defined in ISO/IEC 7816-4:2005(E) 7.5.11
	InstrManageSecurityEnvironment Instruction = "\x22"

	// InstrChangeReferenceData represents the "CHANGE REFERENCE DATA" instruction as defined in ISO/IEC 7816-4:2005(E) 7.5.7
	InstrChangeReferenceData Instruction = "\x24"

	// InstrDisableVerificationRequirement represents the "DISABLE VERIFICATION REQUIREMENT" instruction as defined in ISO/IEC 7816-4:2005(E) 7.5.9
	InstrDisableVerificationRequirement Instruction = "\x26"

	// InstrEnableVerificationRequirement represents the "ENABLE VERIFICATION REQUIREMENT" instruction as defined in ISO/IEC 7816-4:2005(E) 7.5.8
	InstrEnableVerificationRequirement Instruction = "\x28"

	// InstrPerformSecurityOperation represents the "PERFORM SECURITY OPERATION" instruction as defined in ISO/IEC 7816-4:2005(E) Part 8
	InstrPerformSecurityOperation Instruction = "\x2A"

	// InstrResetRetryCounter represents the "RESET RETRY COUNTER" instruction as defined in ISO/IEC 7816-4:2005(E) 7.5.10
	InstrResetRetryCounter Instruction = "\x2C"

	// InstrActivateFile represents the "ACTIVATE FILE" instruction as defined in ISO/IEC 7816-4:2005(E) Part 9
	InstrActivateFile Instruction = "\x44"

	// InstrGenerateAsymmetricKeyPair represents the "GENERATE ASYMMETRIC KEY PAIR" instruction as defined in ISO/IEC 7816-4:2005(E) Part 8
	InstrGenerateAsymmetricKeyPair Instruction = "\x46"

	// InstrManageChannel represents the "MANAGE CHANNEL" instruction as defined in ISO/IEC 7816-4:2005(E) 7.1.2
	InstrManageChannel Instruction = "\x70"

	// InstrExternalAuthenticate represents the "EXTERNAL (/ MUTUAL) AUTHENTICATE" instruction as defined in ISO/IEC 7816-4:2005(E) 7.5.4
	InstrExternalAuthenticate Instruction = "\x82"

	// InstrGetChallenge represents the "GET CHALLENGE" instruction as defined in ISO/IEC 7816-4:2005(E) 7.5.3
	InstrGetChallenge Instruction = "\x84"

	// InstrGeneralAuthenticate represents the "GENERAL AUTHENTICATE" instruction as defined in ISO/IEC 7816-4:2005(E) 7.5.5
	InstrGeneralAuthenticate Instruction = "\x86\x87"

	// InstrInternalAuthenticate represents the "INTERNAL AUTHENTICATE" instruction as defined in ISO/IEC 7816-4:2005(E) 7.5.2
	InstrInternalAuthenticate Instruction = "\x88"

	// InstrSearchBinary represents the "SEARCH BINARY" instruction as defined in ISO/IEC 7816-4:2005(E) 7.2.6
	InstrSearchBinary Instruction = "\xA0\xA1"

	// InstrSearchRecord represents the "SEARCH RECORD" instruction as defined in ISO/IEC 7816-4:2005(E) 7.3.7
	InstrSearchRecord Instruction = "\xA2"

	// InstrSelect represents the "SELECT" instruction as defined in ISO/IEC 7816-4:2005(E) 7.1.1
	InstrSelect Instruction = "\xA4"

	// InstrReadBinary represents the "READ BINARY" instruction as defined in ISO/IEC 7816-4:2005(E) 7.2.3
	InstrReadBinary Instruction = "\xB0\xB1"

	// InstrReadRecord represents the "READ RECORD" instruction as defined in ISO/IEC 7816-4:2005(E) (S) 7.3.3
	InstrReadRecord Instruction = "\xB2\xB3"

	// InstrGetResponse represents the "GET RESPONSE" instruction as defined in ISO/IEC 7816-4:2005(E) 7.6.1
	InstrGetResponse Instruction = "\xC0"

	// InstrEnvelope represents the "ENVELOPE" instruction as defined in ISO/IEC 7816-4:2005(E) 7.6.2
	InstrEnvelope Instruction = "\xC2\xC3"

	// InstrGetData represents the "GET DATA" instruction as defined in ISO/IEC 7816-4:2005(E) 7.4.2
	InstrGetData Instruction = "\xCA\xCB"

	// InstrWriteBinary represents the "WRITE BINARY" instruction as defined in ISO/IEC 7816-4:2005(E) 7.2.6
	InstrWriteBinary Instruction = "\xD0\xD1"

	// InstrWriteRecord represents the "WRITE RECORD" instruction as defined in ISO/IEC 7816-4:2005(E) 7.3.4
	InstrWriteRecord Instruction = "\xD2"

	// InstrUpdateBinary represents the "UPDATE BINARY" instruction as defined in ISO/IEC 7816-4:2005(E) 7.2.5
	InstrUpdateBinary Instruction = "\xD6\xD7"

	// InstrPutData represents the "PUT DATA" instruction as defined in ISO/IEC 7816-4:2005(E) 7.4.3
	InstrPutData Instruction = "\xDA\xDB"

	// InstrUpdateRecord represents the "UPDATE RECORD" instruction as defined in ISO/IEC 7816-4:2005(E) 7.3.5
	InstrUpdateRecord Instruction = "\xDC\xDD"

	// InstrCreateFile represents the "CREATE FILE" instruction as defined in ISO/IEC 7816-4:2005(E) Part 9
	InstrCreateFile Instruction = "\xE0"

	// InstrAppendRecord represents the "APPEND RECORD" instruction as defined in ISO/IEC 7816-4:2005(E) 7.3.6
	InstrAppendRecord Instruction = "\xE2"

	// InstrDeleteFile represents the "DELETE FILE" instruction as defined in ISO/IEC 7816-4:2005(E) Part 9
	InstrDeleteFile Instruction = "\xE4"

	// InstrTerminateDf represents the "TERMINATE DF" instruction as defined in ISO/IEC 7816-4:2005(E) Part 9
	InstrTerminateDf Instruction = "\xE6"

	// InstrTerminateEf represents the "TERMINATE EF" instruction as defined in ISO/IEC 7816-4:2005(E) Part 9
	InstrTerminateEf Instruction = "\xE8"

	// InstrTerminateCardUsage represents the "TERMINATE CARD USAGE" instruction as defined in ISO/IEC 7816-4:2005(E) Part 9
	InstrTerminateCardUsage Instruction = "\xFE"
)

var reverseInstruction = map[Instruction]string{
	"\x04":     "InstrDeactivateFile",
	"\x0C":     "InstrEraseRecord",
	"\x0E\x0F": "InstrEraseBinary",
	"\x10":     "InstrPerformScqlOperation",
	"\x12":     "InstrPerformTransactionOperation",
	"\x14":     "InstrPerformUserOperation",
	"\x20\x21": "InstrVerify",
	"\x22":     "InstrManageSecurityEnvironment",
	"\x24":     "InstrChangeReferenceData",
	"\x26":     "InstrDisableVerificationRequirement",
	"\x28":     "InstrEnableVerificationRequirement",
	"\x2A":     "InstrPerformSecurityOperation",
	"\x2C":     "InstrResetRetryCounter",
	"\x44":     "InstrActivateFile",
	"\x46":     "InstrGenerateAsymmetricKeyPair",
	"\x70":     "InstrManageChannel",
	"\x82":     "InstrExternalAuthenticate",
	"\x84":     "InstrGetChallenge",
	"\x86\x87": "InstrGeneralAuthenticate",
	"\x88":     "InstrInternalAuthenticate",
	"\xA0\xA1": "InstrSearchBinary",
	"\xA2":     "InstrSearchRecord",
	"\xA4":     "InstrSelect",
	"\xB0\xB1": "InstrReadBinary",
	"\xB2\xB3": "InstrReadRecord",
	"\xC0":     "InstrGetResponse",
	"\xC2\xC3": "InstrEnvelope",
	"\xCA\xCB": "InstrGetData",
	"\xD0\xD1": "InstrWriteBinary",
	"\xD2":     "InstrWriteRecord",
	"\xD6\xD7": "InstrUpdateBinary",
	"\xDA\xDB": "InstrPutData",
	"\xDC\xDD": "InstrUpdateRecord",
	"\xE0":     "InstrCreateFile",
	"\xE2":     "InstrAppendRecord",
	"\xE4":     "InstrDeleteFile",
	"\xE6":     "InstrTerminateDf",
	"\xE8":     "InstrTerminateEf",
	"\xFE":     "InstrTerminateCardUsage",
}

func (i Instruction) String() string {
	if v, ok := reverseInstruction[i]; ok {
		return v
	}

	return fmt.Sprintf("unknown instruction %+q", string(i))
}