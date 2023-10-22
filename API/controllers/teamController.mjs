import { ExpressValidator } from "express-validator";
import { tradDb, meanDb, stdDb, zscoreDb } from "../db/connection.mjs";
import orderedTeams from "../util/util.mjs";
import "express-validator"

const { validationResult } = new ExpressValidator

/**
 * 
 * @param {*} req : request object
 * @param {*} res : response object
 * @param {*} next : next function
 * @returns response with either error status and message or success status and data
 */
export const findTwoTeams = async (req, res, next) => {
    try {
        const errors = validationResult(req);

        if (!errors.isEmpty()) {
            return res.status(400).json({
                success: false,
                errors: errors.array(),
            })
        }

        let team1name = req.body.team1;
        let team2name = req.body.team2;
        let year = req.body.year;

        let tradCollection = tradDb.collection("NbaTeamStats_" + year)
        let meanCollection = meanDb.collection("NbaTeamStatsMean_" + year)
        let stdCollection = stdDb.collection("NbaTeamStatsStd_" + year)
        let zscoreCollection = zscoreDb.collection("NbaTeamStatsZscore_" + year)

        // let meanResults = await meanCollection.find({}).toArray();
        // let stdResults = await stdCollection.find({}).toArray();
        

        // let team1 = await tradCollection.find({Name: team1name}).toArray();
        // let team2 = await tradCollection.find({Name: team2name}).toArray();
        // let team1Zscore = await zscoreCollection.find({Name: team1name}).toArray();
        // let team2Zscore = await zscoreCollection.find({Name: team2name}).toArray();

        // Promise.all runs all promises concurrently.
        let data = await Promise.all(
            [
                tradCollection.find({Name: team1name}).toArray(),
                tradCollection.find({Name: team2name}).toArray(),
                meanCollection.find({}).toArray(),
                stdCollection.find({}).toArray(),
                zscoreCollection.find({Name: team1name}).toArray(),
                zscoreCollection.find({Name: team2name}).toArray()
            ]
        )

        // let results = {
        //     team1Traditional: team1,
        //     team2Traditional: team2,
        //     mean: meanResults,
        //     std: stdResults,
        //     team1Zscore: team1Zscore,
        //     team2Zscore: team2Zscore
        // }

        let concurResults = {
            team1Traditional: data[0][0],
            team2Traditional: data[1][0],
            mean: data[2][0],
            std: data[3][0],
            team1Zscore: data[4][0],
            team2Zscore: data[5][0]
        }

        return res.status(200).json(concurResults)
    } catch (err) {
        next(err);
    }
};

export const findTwoTeamsOrdered = async (req, res, next) => {
    try {
        const errors = validationResult(req);

        if (!errors.isEmpty()) {
            return res.status(400).json({
                success: false,
                errors: errors.array(),
            })
        }

        let team1name = req.body.team1;
        let team2name = req.body.team2;
        let year = req.body.year;

        let tradCollection = tradDb.collection("NbaTeamStats_" + year)
        let meanCollection = meanDb.collection("NbaTeamStatsMean_" + year)
        let stdCollection = stdDb.collection("NbaTeamStatsStd_" + year)
        let zscoreCollection = zscoreDb.collection("NbaTeamStatsZscore_" + year)

        let data = await Promise.all(
            [
                tradCollection.find({Name: team1name}).toArray(),
                tradCollection.find({Name: team2name}).toArray(),
                meanCollection.find({}).toArray(),
                stdCollection.find({}).toArray(),
                zscoreCollection.find({Name: team1name}).toArray(),
                zscoreCollection.find({Name: team2name}).toArray()
            ]
        )

        let orderedTeam = orderedTeams(
            data[0][0],
            data[1][0],
            data[2][0],
            data[3][0],
            data[4][0], 
            data[5][0]
        )

        return res.status(200).json(orderedTeam)
    } catch (err) {
        console.log(err)
        next(err);
    }
}

// export default findTwoTeams;