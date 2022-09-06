import logo from './logo.svg';
import './App.css';
import {Nav} from "react-bootstrap";

function App() {
  return (
    <div className="App">
      <header className="App-header">
        <Nav
            activeKey="/home" as="ul"
        >
          <Nav.Item>
            <Nav.Link href="/">Home</Nav.Link>
          </Nav.Item>
          <Nav.Item>
            <Nav.Link href="/rabbitmq/">RabbitMQ</Nav.Link>
          </Nav.Item>
        </Nav>
      </header>
    </div>
  );
}

export default App;
