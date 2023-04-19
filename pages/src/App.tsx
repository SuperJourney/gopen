// 使用react编写一个页面，要求如下，使用bootstrap；需要自适应

// 1. 添加或者删除app ,可以修改属性

// app 有名称和排序

import React from "react";
import { useState } from "react";
import { Container, Row, Col, Button, Form } from "react-bootstrap";

interface AppData {
  name: string;
  order: number;
}

function App() {
  const [apps, setApps] = useState<AppData[]>([
    { name: "App 1", order: 1 },
    { name: "App 2", order: 2 },
    { name: "App 3", order: 3 },
  ]);
  const [newAppName, setNewAppName] = useState<string>("");
  const [newAppOrder, setNewAppOrder] = useState<number>();

  const handleAddApp = () => {
    const newApp: AppData = { name: newAppName, order: newAppOrder! };
    setApps([...apps, newApp]);
    setNewAppName("");
    setNewAppOrder(undefined);
  };

  const handleDeleteApp = (index: number) => {
    const newApps = [...apps];
    newApps.splice(index, 1);
    setApps(newApps);
  };

  const handleEditApp = (index: number, newName: string, newOrder: number) => {
    const newApps = [...apps];
    newApps[index].name = newName;
    newApps[index].order = newOrder;
    setApps(newApps);
  };

  return (
    <Container>
      <Row className="justify-content-center">
        <Col md={6}>
          <h1 className="text-center mb-4">App Manager</h1>
          <Form.Group controlId="formAppName">
            <Form.Label>App Name</Form.Label>
            <Form.Control
              type="text"
              placeholder="Enter app name"
              value={newAppName}
              onChange={(e) => setNewAppName(e.target.value)}
            />
          </Form.Group>
          <Form.Group controlId="formAppOrder">
            <Form.Label>App Order</Form.Label>
            <Form.Control
              type="number"
              placeholder="Enter app order"
              value={newAppOrder}
              onChange={(e) => setNewAppOrder(parseInt(e.target.value))}
            />
          </Form.Group>
          <Button className="mb-3" onClick={handleAddApp}>
            Add App
          </Button>
          <hr />
          {apps.map((app: AppData, index: number) => (
            <div key={index}>
              <Form.Group controlId={`formAppName${index}`}>
                <Form.Label>App Name</Form.Label>
                <Form.Control
                  type="text"
                  defaultValue={app.name}
                  onChange={(e) => {
                    handleEditApp(index, e.target.value, app.order);
                  }}
                />
              </Form.Group>
              <Form.Group controlId={`formAppOrder${index}`}>
                <Form.Label>App Order</Form.Label>
                <Form.Control
                  type="number"
                  defaultValue={app.order}
                  onChange={(e) => {
                    handleEditApp(index, app.name, parseInt(e.target.value));
                  }}
                />
              </Form.Group>
              <Button variant="danger" onClick={() => handleDeleteApp(index)}>
                Delete
              </Button>
              <hr />
            </div>
          ))}
        </Col>
      </Row>
    </Container>
  );
}

export default App;
