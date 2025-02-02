package util

import (
	teamspb "github.com/parathan/NBAMatchups/Microservices/TeamsService/proto"

	"github.com/parathan/NBAMatchups/Microservices/TeamsService/database"
)

// TeamMapping maps a database.TeamData to a teamspb.Team.
//
// It takes a database.TeamData as input and returns a *teamspb.Team and
// an error. It creates a new teamspb.Team by converting the fields of the
// input database.TeamData to float32. The resulting teamspb.Team is then
// returned.
//
// Parameters:
// - team: The database.TeamData to be mapped.
//
// Returns:
// - *teamspb.Team: The mapped teamspb.Team.
// - error: An error if the mapping fails.
func TeamMapping(team database.TeamData) (*teamspb.Team, error) {

	newTeam := &teamspb.Team{
		Id:               team.ID,
		Name:             team.NAME,
		Year: 			  float32(team.YEAR),
		G:                float32(team.G),
		Mp:               float32(team.MP),
		Fg:               float32(team.FG),
		Fga:              float32(team.FGA),
		FgPct:            float32(team.FGPct),
		Fg3:              float32(team.FG3),
		Fg3A:             float32(team.FG3A),
		Fg3Pct:           float32(team.FG3Pct),
		Fg2:              float32(team.FG2),
		Fg2A:             float32(team.FG2A),
		Fg2Pct:           float32(team.FG2Pct),
		Ft:               float32(team.FT),
		Fta:              float32(team.FTA),
		FtPct:            float32(team.FTPct),
		Orb:              float32(team.ORB),
		Drb:              float32(team.DRB),
		Trb:              float32(team.TRB),
		Ast:              float32(team.AST),
		Stl:              float32(team.STL),
		Blk:              float32(team.BLK),
		Tov:              float32(team.TOV),
		Pf:               float32(team.PF),
		Pts:              float32(team.PTS),
		OppFg:            float32(team.OppFG),
		OppFga:           float32(team.OppFGA),
		OppFgPct:         float32(team.OppFGPct),
		OppFg3:           float32(team.OppFG3),
		OppFg3A:          float32(team.OppFG3A),
		OppFg3Pct:        float32(team.OppFG3Pct),
		OppFg2:           float32(team.OppFG2),
		OppFg2A:          float32(team.OppFG2A),
		OppFg2Pct:        float32(team.OppFG2Pct),
		OppFt:            float32(team.OppFT),
		OppFta:           float32(team.OppFTA),
		OppFtPct:         float32(team.OppFTPct),
		OppOrb:           float32(team.OppORB),
		OppDrb:           float32(team.OppDRB),
		OppTrb:           float32(team.OppTRB),
		OppAst:           float32(team.OppAST),
		OppStl:           float32(team.OppSTL),
		OppBlk:           float32(team.OppBLK),
		OppTov:           float32(team.OppTOV),
		OppPf:            float32(team.OppPF),
		OppPts:           float32(team.OppPTS),
		Age:              float32(team.Age),
		Wins:             float32(team.Wins),
		Losses:           float32(team.Losses),
		WinsPyth:         float32(team.WinsPyth),
		LossesPyth:       float32(team.LossesPyth),
		Mov:              float32(team.MOV),
		Sos:              float32(team.SOS),
		Srs:              float32(team.SRS),
		OffRtg:           float32(team.OffRtg),
		DefRtg:           float32(team.DefRtg),
		NetRtg:           float32(team.NetRtg),
		Pace:             float32(team.Pace),
		FtaPerFgaPct:     float32(team.FTAPerFGAPct),
		Fg3APerFgaPct:    float32(team.FG3APerFGAPct),
		TsPct:            float32(team.TSPct),
		EfgPct:           float32(team.EFGPct),
		TovPct:           float32(team.TOVPct),
		OrbPct:           float32(team.ORBPct),
		FtRate:           float32(team.FTRate),
		OppEfgPct:        float32(team.OppEFGPct),
		OppTovPct:        float32(team.OppTOVPct),
		DrbPct:           float32(team.DRBPct),
		OppFtRate:        float32(team.OppFTRate),
		AvgDist:          float32(team.AvgDist),
		PctFgaFg2A:       float32(team.PctFGA2A),
		PctFga_00_03:     float32(team.PctFGA003),
		PctFga_03_10:     float32(team.PctFGA0310),
		PctFga_10_16:     float32(team.PctFGA1016),
		PctFga_16Xx:      float32(team.PctFGA16XX),
		PctFgaFg3A:       float32(team.PctFGA3A),
		FgPct_00_03:      float32(team.FGPct003),
		FgPct_03_10:      float32(team.FGPct0310),
		FgPct_10_16:      float32(team.FGPct1016),
		FgPct_16Xx:       float32(team.FGPct16XX),
		PctAstFg2:        float32(team.PctASTFG2),
		PctAstFg3:        float32(team.PctASTFG3),
		PctFgaDunk:       float32(team.PctFGADunk),
		FgDunk:           float32(team.FGDunk),
		PctFgaLayup:      float32(team.PctFGALayup),
		FgLayup:          float32(team.FGLayup),
		PctFg3ACorner:    float32(team.PctFG3ACorner),
		Fg3PctCorner:     float32(team.FG3PctCorner),
		OppAvgDist:       float32(team.OppAvgDist),
		OppPctFgaFg2A:    float32(team.OppPctFGA2A),
		OppPctFga_00_03:  float32(team.OppPctFGA003),
		OppPctFga_03_10:  float32(team.OppPctFGA0310),
		OppPctFga_10_16:  float32(team.OppPctFGA1016),
		OppPctFga_16Xx:   float32(team.OppPctFGA16XX),
		OppPctFgaFg3A:    float32(team.OppPctFGA3A),
		OppFgPct_00_03:   float32(team.OppFGPct003),
		OppFgPct_03_10:   float32(team.OppFGPct0310),
		OppFgPct_10_16:   float32(team.OppFGPct1016),
		OppFgPct_16Xx:    float32(team.OppFGPct16XX),
		OppPctAstFg2:     float32(team.OppPctASTFG2),
		OppPctAstFg3:     float32(team.OppPctASTFG3),
		OppPctFgaDunk:    float32(team.OppPctFGADunk),
		OppFgDunk:        float32(team.OppFGDunk),
		OppPctFgaLayup:   float32(team.OppPctFGALayup),
		OppFgLayup:       float32(team.OppFGLayup),
		OppPctFg3ACorner: float32(team.OppPctFG3ACorner),
		OppFg3PctCorner:  float32(team.OppFG3PctCorner),
	}

	return newTeam, nil
}