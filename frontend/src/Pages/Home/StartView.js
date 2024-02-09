import React, { useEffect, useContext, useState } from 'react';
import { useNavigate } from 'react-router-dom';
import { HostnamesContext } from "../../Contexts/Hostnames";

const StartView = () => {
    const { hostnames, setHostnames } = useContext(HostnamesContext);
    const navigate = useNavigate();
    const [isLoading, setIsLoading] = useState(false);

    const handleStart = async () => {
        setIsLoading(true);
        if (!localStorage.getItem('sessionID')) {
            await createSession();
        } else {
            await setHostnamesWithExistingSession();
        }
    };

    const createSession = async () => {
        try {
            const response = await fetch('http://labmanager.terminal.kataterm.com:30713/create', {
                method: 'GET'
            });
            const data = await response.json();
            const newUUID = data.newUUID;
            const newHostnames = generateHostnames(newUUID);
            localStorage.setItem('sessionID', newUUID);
            setHostnames(newHostnames);
            await performHealthChecks(newHostnames);
        } catch (error) {
            console.error(error);
        }
    };

    const setHostnamesWithExistingSession = async () => {
        const sessionID = localStorage.getItem('sessionID');
        const newHostnames = generateHostnames(sessionID);
        setHostnames(newHostnames);
        await performHealthChecks(newHostnames);
    };

    const generateHostnames = (sessionID) => {
        const newHostnames = {
            backend: `${sessionID}.backend.terminal.kataterm.com:30713`,
            ttyd: `${sessionID}.ttyd.terminal.kataterm.com:30713`,
            codeServer: `${sessionID}.codeeditor.terminal.kataterm.com:30713`,
        };
        localStorage.setItem('hostnames', JSON.stringify(newHostnames));
        return newHostnames;
    };

    const performHealthChecks = async (hostnames) => {
        if (Object.values(hostnames).some(hostname => !hostname)) {
            console.error("One or more hostnames are missing.");
            if (localStorage.getItem('sessionID')) {
                setHostnamesWithExistingSession(); // Retry setting hostnames
            }
            return;
        }

        const healthChecks = Object.values(hostnames).map(async (hostname) => {
            try {
                const response = await fetch(`http://${hostname}${hostname === hostnames.backend ? '/trials' : '/'}`);
                if (response.status !== 200 && response.status !== 404) {
                    throw new Error(`Health check failed for ${hostname}`);
                }
            } catch (error) {
                console.error(error);
                throw error; // Rethrowing error to be caught in Promise.all
            }
        });

        try {
            await Promise.all(healthChecks);
            navigate('/quiz');
        } catch (error) {
            console.error("Health check failed for one or more services.");
            setTimeout(() => performHealthChecks(hostnames), 20000);
        } finally {
            setIsLoading(false);
        }
    };

    useEffect(() => {
        if (Object.keys(hostnames).length > 0) {
            const interval = setInterval(() => {
                performHealthChecks(hostnames);
            }, 20000);

            return () => clearInterval(interval);
        }
    }, [hostnames]);

    return (
        <div>
            <h1>Welcome to Quizard!</h1>
            <button onClick={handleStart} disabled={isLoading}>Start</button>
        </div>
    );
};

export default StartView;