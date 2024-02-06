import React, { useEffect, useContext } from 'react';
import { useNavigate } from 'react-router-dom';
import { HostnamesContext } from "../../Contexts/Hostnames";

const StartView = () => {

    const { hostnames, setHostnames } = useContext(HostnamesContext)
    const navigate = useNavigate();


    const handleStart = () => {
        // Hacer la primera solicitud al servicio proxy
        fetch('http://terminal.kataterm.com:30009/create')
            .then(response => response.json())
            .then(data => {
                // Supongamos que el servicio proxy devuelve el nuevo UUID del backend
                const newUUID = data.newUUID;

                // Construir los nuevos hostnames utilizando el nuevo UUID
                const newHostnames = {
                    backend: `${newUUID}.backend.terminal.kataterm.com:30007`,
                    ttyd: `${newUUID}.ttyd.terminal.kataterm.com:30007`,
                    codeServer: `${newUUID}.codeeditor.terminal.kataterm.com:30007`,
                };

                // Usage example
                console.log(`Backend endpoint: http://${newHostnames.backend}`);
                console.log(`TTYD endpoint: http://${newHostnames.ttyd}`);
                console.log(`Code-server endpoint: http://${newHostnames.codeServer}`);
                // Actualizar los hostnames utilizando el nuevo objeto de hostnames
                setHostnames(newHostnames);

                // Navegar a la vista del quiz
                navigate('/quiz');
            });
    };

    return (
        <div>
            <h1>Welcome to Quizard!</h1>
            <button onClick={handleStart}>Start</button>
        </div>
    );
};

export default StartView;
