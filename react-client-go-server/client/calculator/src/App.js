import React from 'react';
import { CalculatorClient } from './calculator_grpc_web_pb';
import { AddRequest, SubtractRequest } from './calculator_pb';

const client = new CalculatorClient('http://localhost:8000');

function App() {
  const handleAdd = () => {
    const request = new AddRequest();
    request.setOperand1(10);
    request.setOperand2(5);

    client.add(request, {}, (err, response) => {
      if (err) {
        console.error('Addition failed:', err);
        return;
      }
      console.log('Addition result:', response.getResult());
    });
  };

  const handleSubtract = () => {
    const request = new SubtractRequest();
    request.setOperand1(10);
    request.setOperand2(5);

    client.subtract(request, {}, (err, response) => {
      if (err) {
        console.error('Subtraction failed:', err);
        return;
      }
      console.log('Subtraction result:', response.getResult());
    });
  };

  return (
      <div>
        <button onClick={handleAdd}>Add</button>
        <button onClick={handleSubtract}>Subtract</button>
      </div>
  );
}

export default App;