import React, { useState } from "react";
import { Form, Button } from "react-bootstrap";
import { useDispatch } from "react-redux";

import { registerUser } from "../redux/user/register/userRegisterAction.js";

const Register = () => {
  const dispatch = useDispatch();

  const [fullname, setFullname] = useState("");
  const [address, setAddress] = useState("");
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");

  const registerSubmit = () => {
    const data = {
      fullname: fullname,
      address: address,
      email: email,
      password: password,
    };

    console.log(data);

    dispatch(registerUser(data));
  };

  return (
    <div>
      <h1>Register</h1>
      <Form style={{ margin: "50px" }} onSubmit={registerSubmit}>
        <Form.Group className="mb-3" controlId="formBasicEmail">
          <Form.Label>Fullname</Form.Label>
          <Form.Control
            type="text"
            className="form-control"
            id="fullname"
            onChange={(e) => {
              e.preventDefault();
              setFullname(e.target.value);
            }}
          />
        </Form.Group>
        <Form.Group className="mb-3" controlId="formBasicEmail">
          <Form.Label>Address</Form.Label>
          <Form.Control
            type="text"
            placeholder="Enter Address"
            id="address"
            onChange={(e) => {
              e.preventDefault();
              setAddress(e.target.value);
            }}
          />
        </Form.Group>
        <Form.Group className="mb-3" controlId="formBasicEmail">
          <Form.Label>Email address</Form.Label>
          <Form.Control
            type="email"
            placeholder="Enter email"
            id="email"
            onChange={(e) => {
              e.preventDefault();
              setEmail(e.target.value);
            }}
          />
        </Form.Group>
        <Form.Group className="mb-3" controlId="formBasicPassword">
          <Form.Label>Password</Form.Label>
          <Form.Control
            type="password"
            placeholder="Password"
            id="password"
            onChange={(e) => {
              e.preventDefault();
              setPassword(e.target.value);
            }}
          />
        </Form.Group>
        <Form.Group className="mb-3" controlId="formBasicCheckbox">
          <Form.Check type="checkbox" label="Check me out" />
        </Form.Group>
        <Button variant="primary" type="submit">
          Submit
        </Button>
      </Form>
    </div>
  );
};

export default Register;
