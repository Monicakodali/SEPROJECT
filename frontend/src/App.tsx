import React from 'react';
import { Layout } from './components'
import { ThemeProvider } from '@mui/material/styles';
import theme from './theme'
import {
  BrowserRouter as Router,
  Routes,
  Route
} from "react-router-dom";
import { LandingPage, Locations, Establishment } from './pages'

function App() {
  return (
    <ThemeProvider theme={theme}>
      <Router>
        <Routes>
          <Route path="/" element={<Layout header={false}><LandingPage /></Layout>} />
          <Route path="/search" element={<Layout><Locations /></Layout>} />
          <Route path="/biz/:id" element={<Layout><Establishment /></Layout>} />
        </Routes>
      </Router>
    </ThemeProvider>
  );
}

export default App;
