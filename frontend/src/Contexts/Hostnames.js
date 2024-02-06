import React, { createContext, useState } from 'react';

export const HostnamesContext = createContext();

export const HostnamesProvider = ({ children }) => {
    const [hostnames, setHostnames] = useState({
        backend: '',
        ttyd: '',
        codeServer: ''
    });


    return (
        <HostnamesContext.Provider value={{ hostnames, setHostnames }}>
            {children}
        </HostnamesContext.Provider>
    );
};