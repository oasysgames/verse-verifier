package database

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

func TestOptimizeDatabase(t *testing.T) {
	suite.Run(t, new(OptimizeDatabaseTestSuite))
}

type OptimizeDatabaseTestSuite struct {
	DatabaseTestSuite
}

func (s *OptimizeDatabaseTestSuite) TestRepairOvertakingSignatures() {
	signer := s.createSigner()
	scc := s.createSCC()
	ids := []string{
		"01GSSK5XTCZD5ZCXR5E487F1Q1", //  0
		"01GSSK5XTDBST5SJZ68JRX35CV", //  1
		"01GSSK5XTFNZ8P82AHZ2VCR2Y0", //  2
		"01GSSK5XTGB76GRS6Q88PV5YY2", //  3
		"01GSSK5XTHHWYPYS0C9X1501TQ", //  4 (overtaking)
		"01GSSK5XTJA4CB36RDRR00D6AZ", //  5
		"01GSSK5XTKKKY8YDFXK83JP80A", //  6
		"01GSSK5XTM5YTTK0997R6R1XY9", //  7
		"01GSSK5XTPTQ35ZEHAM0F2NHNM", //  8
		"01GSSK5XTQC80MA6KNY0QJRX0W", //  9
		"01GSSK5XTR74VEM1FAC5RZS30Q", // 10
		"01GSSK5XTSZPD1A01C8N7287EP", // 11
		"01GSSK5XTTQNZX0V8790TE9HQT", // 12
		"01GSSK5XTWKSMTMGHXRKG8SHZB", // 13
		"01GSSK5XTX62C2N97TD0F3PXDR", // 14
		"01GSSK5XTYJJMR02T9HH13Z1F8", // 15
		"01GSSK5XTZB6EQK946EMQ2SBAE", // 16
		"01GSSK5XV05HN6M7TFFGY7N3HH", // 17 (overtaking)
		"01GSSK5XV282QQQQF91KGQJRY9", // 18
		"01GSSK5XV3M6D757K6G43H32A6", // 19
	}

	for i, id := range ids {
		prevID := ""
		if i > 0 {
			prevID = ids[i-1]
		}

		// overtaking
		if i == 4 || i == 17 {
			prevID = ids[i+1]
		}

		sig := &OptimismSignature{
			ID:          id,
			PreviousID:  prevID,
			Signer:      *signer,
			OptimismScc: *scc,
			BatchIndex:  uint64(i),
			BatchRoot:   s.RandHash(),
			Signature:   RandSignature(),
		}
		s.NoDBError(s.rawdb.Create(sig))
	}

	var rows []*OptimismSignature
	tx := s.rawdb.Where("id < previous_id").Order("id")

	// check if overtaking
	tx.Find(&rows)
	s.Len(rows, 2)

	actual1, actual2 := rows[0], rows[1]

	s.Equal(actual1.ID, "01GSSK5XTHHWYPYS0C9X1501TQ")
	s.Equal(actual1.PreviousID, "01GSSK5XTJA4CB36RDRR00D6AZ")
	s.Less(actual1.ID, actual1.PreviousID)

	s.Equal(actual2.ID, "01GSSK5XV05HN6M7TFFGY7N3HH")
	s.Equal(actual2.PreviousID, "01GSSK5XV282QQQQF91KGQJRY9")
	s.Less(actual2.ID, actual2.PreviousID)

	// run repair
	s.db.Optimism.repairOvertakingSignatures(signer.Address)

	// check if repaired
	tx.Find(&rows)
	s.Len(rows, 0)

	actual1, _ = s.db.Optimism.FindSignatureByID(actual1.ID)
	actual2, _ = s.db.Optimism.FindSignatureByID(actual2.ID)

	s.Equal(actual1.PreviousID, "01GSSK5XTGB76GRS6Q88PV5YY2")
	s.Greater(actual1.ID, actual1.PreviousID)

	s.Equal(actual2.PreviousID, "01GSSK5XTZB6EQK946EMQ2SBAE")
	s.Greater(actual2.ID, actual2.PreviousID)
}

func (s *OptimizeDatabaseTestSuite) TestRepairMissingPrevID() {
	signer := s.createSigner()
	scc := s.createSCC()
	ids := []string{
		"01GSSK5XTCZD5ZCXR5E487F1Q1", //  0
		"01GSSK5XTDBST5SJZ68JRX35CV", //  1
		"01GSSK5XTFNZ8P82AHZ2VCR2Y0", //  2
		"01GSSK5XTGB76GRS6Q88PV5YY2", //  3
		"01GSSK5XTHHWYPYS0C9X1501TQ", //  4 (missing)
		"01GSSK5XTJA4CB36RDRR00D6AZ", //  5
		"01GSSK5XTKKKY8YDFXK83JP80A", //  6
		"01GSSK5XTM5YTTK0997R6R1XY9", //  7
		"01GSSK5XTPTQ35ZEHAM0F2NHNM", //  8
		"01GSSK5XTQC80MA6KNY0QJRX0W", //  9
		"01GSSK5XTR74VEM1FAC5RZS30Q", // 10
		"01GSSK5XTSZPD1A01C8N7287EP", // 11
		"01GSSK5XTTQNZX0V8790TE9HQT", // 12
		"01GSSK5XTWKSMTMGHXRKG8SHZB", // 13
		"01GSSK5XTX62C2N97TD0F3PXDR", // 14
		"01GSSK5XTYJJMR02T9HH13Z1F8", // 15
		"01GSSK5XTZB6EQK946EMQ2SBAE", // 16
		"01GSSK5XV05HN6M7TFFGY7N3HH", // 17 (missing)
		"01GSSK5XV282QQQQF91KGQJRY9", // 18
		"01GSSK5XV3M6D757K6G43H32A6", // 19
	}

	for i, id := range ids {
		prevID := ""
		if i > 0 {
			prevID = ids[i-1]
		}

		// missing
		if i == 4 {
			prevID = "01GSXXVYJ2K3DZWC8FYYMTEP0D"
		} else if i == 17 {
			prevID = "01GT065ZW542460G75SPMJ2BC0"
		}

		sig := &OptimismSignature{
			ID:          id,
			PreviousID:  prevID,
			Signer:      *signer,
			OptimismScc: *scc,
			BatchIndex:  uint64(i),
			BatchRoot:   s.RandHash(),
			Signature:   RandSignature(),
		}
		s.NoDBError(s.rawdb.Create(sig))
	}

	var rows []*OptimismSignature
	sub := s.rawdb.Model(&OptimismSignature{}).Select("id")
	tx := s.rawdb.
		Where("previous_id != ''").
		Where("previous_id NOT IN (?)", sub).
		Order("id")

	// check if overtaking
	tx.Find(&rows)
	s.Len(rows, 2)

	actual1, actual2 := rows[0], rows[1]

	s.Equal(actual1.ID, "01GSSK5XTHHWYPYS0C9X1501TQ")
	s.Equal(actual1.PreviousID, "01GSXXVYJ2K3DZWC8FYYMTEP0D")

	s.Equal(actual2.ID, "01GSSK5XV05HN6M7TFFGY7N3HH")
	s.Equal(actual2.PreviousID, "01GT065ZW542460G75SPMJ2BC0")

	// run repair
	s.db.Optimism.repairMissingPrevID(signer.Address)

	// check if repaired
	tx.Find(&rows)
	s.Len(rows, 0)

	actual1, _ = s.db.Optimism.FindSignatureByID(actual1.ID)
	actual2, _ = s.db.Optimism.FindSignatureByID(actual2.ID)

	s.Equal(actual1.PreviousID, "01GSSK5XTGB76GRS6Q88PV5YY2")
	s.Greater(actual1.ID, actual1.PreviousID)

	s.Equal(actual2.PreviousID, "01GSSK5XTZB6EQK946EMQ2SBAE")
	s.Greater(actual2.ID, actual2.PreviousID)
}
