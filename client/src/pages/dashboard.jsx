import React from "react";
import { Table, Button, Form } from "react-bootstrap";

const Dashboard = () => {
  return (
    <div style={{ margin: "50px" }}>
      <Form style={{ marginBottom: "20px" }}>
        <Form.Group className="mb-3" controlId="exampleForm.ControlInput1">
          <Form.Label>Website</Form.Label>
          <Form.Control type="text" placeholder="add list website" />
        </Form.Group>
        <Form.Group className="mb-3" controlId="exampleForm.ControlInput1">
          <Form.Label>Password</Form.Label>
          <Form.Control type="password" placeholder="your password" />
        </Form.Group>
        <Button variant="primary" type="submit">
          Add List
        </Button>
      </Form>

      <Table striped bordered hover variant="dark">
        <thead>
          <tr>
            <th>ID List</th>
            <th>ID User</th>
            <th>Website</th>
            <th>Password</th>
            <th>Created At</th>
            <th>Updated At</th>
            <th></th>
          </tr>
        </thead>
        <tbody>
          <tr>
            <td>1</td>
            <td>Mark</td>
            <td>Otto</td>
            <td>@mdo</td>
            <td>Otto</td>
            <td>@mdo</td>
            <Button variant="primary">Edit</Button>
            <Button variant="secondary">Delete</Button>
          </tr>
        </tbody>
      </Table>
    </div>
  );
};

export default Dashboard;
