import React from 'react';
import { Layout } from './components'
import { ThemeProvider } from '@mui/material/styles';
import theme from './theme'
import {
  BrowserRouter as Router,
  Routes, Route, 
} from "react-router-dom";
import { Locations } from './pages'

function App() {
  return (
    <ThemeProvider theme={theme}>
      <Router>
        <Layout>
          <Routes>
            <Route path="/" element={<Locations />} />
          </Routes>
        </Layout>
      </Router>
    </ThemeProvider>
  );
}

export default App;
