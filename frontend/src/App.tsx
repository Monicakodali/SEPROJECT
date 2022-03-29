import React from 'react';
import { Layout } from './components'
import { ThemeProvider } from '@mui/material/styles';
import theme from './theme'
import {
  BrowserRouter as Router,
  Routes,
  Route
} from "react-router-dom";
import { LandingPage, Locations, Establishment, LoginPage, SignupPage } from './pages'
import { CypressHistorySupport } from 'cypress-react-router';

function App() {
  
  return (
    <ThemeProvider theme={theme}>
      <Router>
        <CypressHistorySupport />
        <Routes>
          <Route path="/" element={<Layout header={false}><LandingPage /></Layout>} />
          <Route path="/search" element={<Layout><Locations /></Layout>} />
          <Route path="/login" element={<Layout header={false}><LoginPage /></Layout>} />
          <Route path="/signup" element={<Layout header={false}><SignupPage /></Layout>} />
          <Route path="/est/:id" element={<Layout><Establishment /></Layout>} />
        </Routes>
      </Router>
    </ThemeProvider>
  );
}

export default App;
