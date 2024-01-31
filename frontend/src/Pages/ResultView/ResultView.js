import React, { useEffect, useState } from 'react';
import './ResultView.css';

const ResultView = () => {
    const [scores, setScores] = useState(null);

    useEffect(() => {
        fetch('http://terminal.kataterm.com:8000/get_score')
            .then(response => response.json())
            .then(data => setScores(data))
            .catch(error => console.log(error));
    }, []);

    const calculatePercentage = (score) => {
        return score.toFixed(3);
    };

    return (
        <div className="result-view">
            <h2 className="result-view__title">Results</h2>
            {scores && (
                <div className="result-view__scores">
                    <p className="result-view__score">Global Score: {calculatePercentage(scores.Global_Score)}%</p>
                    <p className="result-view__score">Config Test Score: {calculatePercentage(scores.Config_Test_Score)}%</p>
                    <p className="result-view__score">Multiple Choice Score: {calculatePercentage(scores.Multiple_Choice_Score)}%</p>
                </div>
            )}
        </div>
    );
};

export default ResultView;