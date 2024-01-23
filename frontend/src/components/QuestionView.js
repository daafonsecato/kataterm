import React, { useState, useRef } from 'react';
import TaskPanel from './TaskPanel';
import TerminalPanel from './TerminalPanel';
import Split from 'react-split';
import EditorPanel from './EditorPanel';

function QuestionView() {
    const ttydUrl = process.env.REACT_APP_TTYD_SERVER_URL; // Replace with your ttyd URL
    const codeServerUrl = process.env.REACT_APP_CODE_SERVER_URL; // Replace with your ttyd URL
    const [isModalOpen, setIsModalOpen] = useState(false);
    const [validationResult, setValidationResult] = useState(null);

    const toggleModal = () => {
        setIsModalOpen(prevState => !prevState); // Toggle the state
    };

    const sendCommandToTerminal = useRef(null);

    const handleConfirmClick = () => {
        console.log('Sending command to terminal...');
        // Send the validation command to the terminal
        if (sendCommandToTerminal.current) {
            console.log('Sending command to terminal iinner...');
            sendCommandToTerminal.current('./public/tasks/task1/validate_task1.sh');
        }

        // Close the modal
        toggleModal();
    };

    const handleValidationResult = (success) => {
        if (success) {
            console.log('Task completed successfully.');
            // Handle success scenario
        } else {
            console.log('Task validation failed.');
            // Handle failure scenario
        }
        setValidationResult(success);
    };

    return (
        <div className="QuestionView">
            <div className="BasePanels">
                <Split
                    className="App"
                    direction="horizontal" // New split direction
                    sizes={[30, 70]} // Initial sizes of the three panels (left, middle, right)
                    minSize={100} // Minimum size of each panel
                    expandToMin={false} // Prevent panels from collapsing to minimum size
                    gutterSize={10} // Size of the gutter between panels
                    gutterAlign="center" // Align the gutter in the center
                    snapOffset={30} // Snap to minimum                                                   size if within 30 pixels
                >
                    <div className="LeftPanel">
                        <TaskPanel onCheckClick={toggleModal} />
                    </div>

                    <Split
                        className="RightPanel"
                        direction="vertical" // New split direction
                        sizes={[60, 40]} // Initial sizes of the two panels (top, bottom)
                        minSize={10} // Minimum size of each panel
                        expandToMin={false} // Prevent panels from collapsing to minimum size
                        gutterSize={10} // Size of the gutter between panels
                        gutterAlign="center" // Align the gutter in the center
                        snapOffset={30} // Snap to minimum size if within 30 pixels
                    >
                        <div className="TopPanel">
                            <EditorPanel codeServerUrl={codeServerUrl} />
                        </div>
                        <div className="BottomPanel">
                            <TerminalPanel wsUrl={ttydUrl} onCommand={(cmd) => sendCommandToTerminal.current = cmd} onValidationResult={handleValidationResult} />
                        </div>
                    </Split>
                </Split>
            </div>

            {/* Modal */}
            {isModalOpen && (
                <div className="modal-background">
                    <div className="flex justify-start content-center" style={{ marginTop: '50px' }}>
                        <div>
                            <div id="modal-1702509518071" className="modal" data-testid="config-test-kke-modal" tabIndex="0">
                                <div className="modal-content">
                                    <h4>Finish?</h4>
                                    <p> Are you sure you want to mark the task completed? You must verify your work to make sure
                                        it has been completed as expected. </p><br />
                                    <p> Once marked finished, you will not be able to make any changes. </p>
                                </div>
                                <div className="modal-footer">
                                    <button className="waves-effect waves-green btn-flat" onClick={toggleModal}>Cancel</button>
                                    <button className="waves-effect waves-green btn-flat" onClick={handleConfirmClick}>Confirm</button>
                                </div>
                            </div>
                        </div>
                    </div>
                </div>
            )}
        </div>
    );
}

export default QuestionView;