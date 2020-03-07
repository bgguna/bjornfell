import React, {
    Component
} from 'react';
import axios from 'axios';
import Form from 'react-bootstrap/Form';

class Contact extends Component {

    constructor(props) {
        super(props);
        this.state = {
            name: '',
            email: '',
            message: ''
        }
    }

    handleSubmit(event) {
        event.preventDefault();
        axios({
            method: 'POST',
            url: 'http://localhost:3002/send',
            data: this.state
        }).then((response) => {
            if (response.data.status === 'success') {
                alert('Message sent.');
                this.resetForm()
            } else if (response.data.status === 'fail') {
                alert('Message failed to send.')
            }
        })
    }

    resetForm() {
        this.setState({
            name: '',
            email: '',
            message: ''
        })
    }

    render() {
        return (
            <div>
                <div>
                    <h2>get in touch?</h2>
                </div>

                <div>
                    <Form id='contact-form' onSubmit={this.handleSubmit.bind(this)} method="POST">
                        <Form.Group controlId="name">
                            <Form.Label>name</Form.Label>
                            <Form.Control type="text" placeholder="enter name" value={this.state.name} onChange={this.onNameChange.bind(this)} />
                        </Form.Group>
                        <Form.Group controlId="email">
                            <Form.Label>email</Form.Label>
                            <Form.Control type="email" placeholder="enter email" value={this.state.email} onChange={this.onEmailChange.bind(this)} />
                        </Form.Group>
                        <Form.Group controlId="message">
                            <Form.Label>message</Form.Label>
                            <Form.Control type="text" placeholder="enter message" value={this.state.message} onChange={this.onMessageChange.bind(this)} />
                        </Form.Group>

                        <button type='submit' className='btn btn-primary'>submit</button>
                    </Form>
                </div>
            </div>
        );
    }

    onNameChange(event) {
        this.setState({
            name: event.target.value
        })
    }

    onEmailChange(event) {
        this.setState({
            email: event.target.value
        })
    }

    onMessageChange(event) {
        this.setState({
            message: event.target.value
        })
    }
}

export default Contact;
