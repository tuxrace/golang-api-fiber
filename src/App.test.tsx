import React from 'react';
import { render } from '@testing-library/react';
import App from './App';

describe('App', () => {
  test('App render', () => {
    const { container } = render(<App />);
    expect(container).toBeDefined();
  });
  
})
