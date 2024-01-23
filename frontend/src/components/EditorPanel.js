import React from 'react';

const EditorPanel = ({ codeServerUrl }) => {


    return (
        <div className='simple-terminal-panel'>
            <iframe className='simple-editor-panel'
                src={codeServerUrl}
                title="VSCode"
            />
        </div>
    );
};

export default EditorPanel;
