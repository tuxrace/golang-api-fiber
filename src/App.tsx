import React from 'react';
import './App.css';
import AppliancesList from './components/AppliancesList';
import { Container } from '@material-ui/core';

function App() {
  return (
    <Container maxWidth="lg">
      <AppliancesList />
    </Container>
  );
}

export default App;
