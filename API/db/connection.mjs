import { MongoClient } from "mongodb";

const connectionString = process.env.MONGO_URI || "";

const client = new MongoClient(connectionString);

let connection;
try {
    connection = await client.connect();
} catch(e) {
    console.error(e)
}

let tradDb = connection.db("NBAMatchups")
let meanDb = connection.db("NBAMatchupsMean")
let stdDb = connection.db("NBAMatchupsStd")
let zscoreDb = connection.db("NBAMatchupsZscore")

export { tradDb, meanDb, stdDb, zscoreDb};