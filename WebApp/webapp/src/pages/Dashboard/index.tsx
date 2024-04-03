import React, { ChangeEvent, useEffect, useState } from 'react';
import styles from './index.module.css';
import Layout from "../../components/Layout/Layout";
import { CircularProgress, Grid } from '@mui/material';

import axios from 'axios';
import { totalTeamData } from '../../interfaces/TotalTeamData';


function Dashboard() {

    const [data, setData] = useState<totalTeamData[]>([])
    const [progress, setProgress] = useState(true)
    const [error, setError] = useState(false)
    const [startYear, setStartYear] = useState(2019)
    const [endYear, setEndYear] = useState(2023)
    const [years, setYears] = useState<string[]>([])

    // Intended that useEffect runs once, as empty array given means that it only updates once when page renders.
    // In development, it will run twice due to strict mode being on. This won't happen in production, however.
    // https://byby.dev/useeffect-run-twice
    useEffect(() => {
        axios.post(process.env.REACT_APP_API_URL + '/teams/allTeams', {
            startYear: startYear,
            endYear: endYear
        })
        .then((response) => {
            console.log(response)
            for (let i = startYear; i <= endYear; i++) {
                setYears(oldArray => [...oldArray, i.toString()])
            }
            setData(response.data)
            setProgress(false)
        })
        .catch((error) => {
            console.log(error)
            setProgress(false)
            setError(true)
        })
    }, [])


    return (
        <Layout>
            <div className={styles.header}>Dashboard</div>
            <div className={`${progress ? styles.notHidden : styles.hidden}`}>
                <CircularProgress />
            </div>
            <div className={`${error ? styles.notHidden : styles.hidden}`}>
                Errors existing
            </div>
        </Layout>
    )
}

export default Dashboard;