export interface TeamData {
    _id: string,
    g: number,
    mp: number,
    fg: number,
    fga: number,
    fg_pct: number,
    fg3: number,
    fg3a: number,
    fg3_pct: number,
    fg2: number,
    fg2a: number,
    fg2_pct: number,
    ft: number,
    fta: number,
    ft_pct: number,
    orb: number,
    drb: number,
    trb: number,
    ast: number,
    stl: number,
    blk: number,
    tov: number,
    pf: number,
    pts: number,
    opp_fg: number,
    opp_fga: number,
    opp_fg_pct: number,
    opp_fg3: number,
    opp_fg3a: number,
    opp_fg3_pct: number,
    opp_fg2: number,
    opp_fg2a: number,
    opp_fg2_pct: number,
    opp_ft: number,
    opp_fta: number,
    opp_ft_pct: number,
    opp_orb: number,
    opp_drb: number,
    opp_trb: number,
    opp_ast: number,
    opp_stl: number,
    opp_blk: number,
    opp_tov: number,
    opp_pf: number,
    opp_pts: number,
    age: number,
    wins: number,
    losses: number,
    wins_pyth: number,
    losses_pyth: number,
    mov: number,
    sos: number,
    srs: number,
    off_rtg: number,
    def_rtg: number,
    net_rtg: number,
    pace: number,
    fta_per_fga_pct: number,
    fg3a_per_fga_pct: number,
    ts_pct: number,
    efg_pct: number,
    tov_pct: number,
    orb_pct: number,
    ft_rate: number,
    opp_efg_pct: number,
    opp_tov_pct: number,
    drb_pct: number,
    opp_ft_rate: number,
    avg_dist: number,
    pct_fga_fg2a: number,
    pct_fga_00_03: number,
    pct_fga_03_10: number,
    pct_fga_10_16: number,
    pct_fga_16_xx: number,
    pct_fga_fg3a: number,
    fg_pct_00_03: number,
    fg_pct_03_10: number,
    fg_pct_10_16: number,
    fg_pct_16_xx: number,
    pct_ast_fg2: number,
    pct_ast_fg3: number,
    pct_fga_dunk: number,
    fg_dunk: number,
    pct_fga_layup: number,
    fg_layup: number,
    pct_fg3a_corner: number,
    fg3_pct_corner: number,
    opp_avg_dist: number,
    opp_pct_fga_fg2a: number,
    opp_pct_fga_00_03: number,
    opp_pct_fga_03_10: number,
    opp_pct_fga_10_16: number,
    opp_pct_fga_16_xx: number,
    opp_pct_fga_fg3a: number,
    opp_fg_pct_00_03: number,
    opp_fg_pct_03_10: number,
    opp_fg_pct_10_16: number,
    opp_fg_pct_16_xx: number,
    opp_pct_ast_fg2: number,
    opp_pct_ast_fg3: number,
    opp_pct_fga_dunk: number,
    opp_fg_dunk: number,
    opp_pct_fga_layup: number,
    opp_fg_layup: number,
    opp_pct_fg3a_corner: number,
    opp_fg3_pct_corner: number,
}

interface YearlyTeamData {
    year: string,
    yearStats: TeamData
}

export interface TotalTeamData {
    teamName: string,
    stats: YearlyTeamData[]
}