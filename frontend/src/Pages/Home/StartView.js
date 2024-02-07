import React, { useEffect, useContext, useState } from 'react';
import { useNavigate } from 'react-router-dom';
import { HostnamesContext } from "../../Contexts/Hostnames";

const StartView = () => {

    const { hostnames, setHostnames } = useContext(HostnamesContext)
    const navigate = useNavigate();
    const [isLoading, setIsLoading] = useState(false);

    const handleStart = () => {
        setIsLoading(true);
        if (!localStorage.getItem('sessionID')) {
            createSession();
        } else {
            setHostnamesWithExistingSession();
        }
    };

    const createSession = () => {
        fetch('http://terminal.kataterm.com:30009/create')
            .then(response => response.json())
            .then(data => {
                const newUUID = data.newUUID;
                const newHostnames = generateHostnames(newUUID);
                localStorage.setItem('sessionID', newUUID);
                setHostnames(newHostnames);
                return newHostnames;
            })
            .then(newHostnames => {
                performHealthChecks(newHostnames);
            })
            .catch(error => {
                console.error(error);
            });
    };

    const setHostnamesWithExistingSession = () => {
        const sessionID = localStorage.getItem('sessionID');
        const newHostnames = generateHostnames(sessionID);
        setHostnames(newHostnames);
        performHealthChecks(newHostnames);
    };

    const generateHostnames = (sessionID) => {
        let hostnames = {
            backend: `${sessionID}.backend.terminal.kataterm.com:30007`,
            ttyd: `${sessionID}.ttyd.terminal.kataterm.com:30007`,
            codeServer: `${sessionID}.codeeditor.terminal.kataterm.com:30007`,
        }
        console.log(hostnames);
        return hostnames;
    };

    const performHealthChecks = (hostnames) => {
        const healthChecks = Object.values(hostnames).map(hostname =>
            fetch(`http://${hostname}${hostname === hostnames.backend ? '/trials' : '/'}`)
                .then(response => {
                    if (response.status === 200 || response.status === 404) {
                        return Promise.resolve();
                    } else {
                        return Promise.reject(new Error(`Health check failed for ${hostname}`));
                    }
                })
                .catch(error => {
                    console.error(error);
                    return Promise.reject(new Error(`Health check failed for ${hostname}`));
                })
        );

        const checkHealth = () => {
            Promise.all(healthChecks)
                .then(() => {
                    navigate('/quiz');
                })
                .catch(error => {
                    console.error(error);
                    setTimeout(checkHealth, 20000);
                })
                .finally(() => {
                    setIsLoading(false);
                });
        };

        checkHealth();
    };

    useEffect(() => {
        const interval = setInterval(() => {
            performHealthChecks(hostnames);
        }, 20000);

        return () => clearInterval(interval);
    }, []);

    return (
        <div>
            <h1>Welcome to Quizard!</h1>
            <button onClick={handleStart} disabled={isLoading}>Start</button>
        </div>
    );
};

export default StartView;
