import React from 'react';
import 'xterm/css/xterm.css';
import './TerminalPanel.css';

function TerminalPanel({ wsUrl }) { // Make sure to receive onCommand as a prop


    return (
        <iframe className="simple-terminal-panel" src={wsUrl} title="Terminal" />
    );
}

export default TerminalPanel;
