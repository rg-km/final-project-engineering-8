import { Modal } from 'react-bootstrap';
import { Button } from 'react-bootstrap';

function CenteredModal(props) {
    return (
        <Modal
            {...props}
            size="md"
            aria-labelledby="contained-modal-title-vcenter"
            centered
        >

            <Modal.Body>
                <h4>Informasi</h4>
                <p>
                    {props.message}
                </p>
            </Modal.Body>
            <Modal.Footer>
                <Button onClick={props.onHide}>Close</Button>
            </Modal.Footer>
        </Modal>
    );
}

export default CenteredModal;