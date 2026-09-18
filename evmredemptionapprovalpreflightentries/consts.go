package evmredemptionapprovalpreflightentries

import "github.com/pkg/errors"

type CoverageT string

const (
	CoverageUnrepresented      CoverageT = "unrepresented"
	CoverageFulfillmentCovered CoverageT = "fulfillment_covered"
	CoverageClaimCovered       CoverageT = "claim_covered"
)

func AllCoverageT() []CoverageT {
	return []CoverageT{CoverageUnrepresented, CoverageFulfillmentCovered, CoverageClaimCovered}
}
func (value CoverageT) IsValid() error {
	for _, candidate := range AllCoverageT() {
		if value == candidate {
			return nil
		}
	}
	return errors.New("enum is not valid")
}
func (value CoverageT) String() string { return string(value) }
