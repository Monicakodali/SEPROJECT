import { red } from '@mui/material/colors';
import { createTheme } from '@mui/material/styles';

const theme = createTheme({
  palette: {
    primary: {
      main: red[500],
    },
  },
  mixins: {
    toolbar: {
      '@media (min-width:0px) and (orientation: landscape)': {
        minHeight: 64
      },
      '@media (min-width:600px)': {
        minHeight: 88
      },
      minHeight: 88
    }
  },
  components: {
    MuiButton: {
      styleOverrides: {
        root: {
          textTransform: 'capitalize',
        },
      },
    },
  },
  typography: {
    h4: {
      fontSize: 24,
      fontWeight: 700,
      lineHeight: 1.1,
      letterSpacing: -0.25,
      color: 'rgba(45,46,47,1)'
    }
  }
});

export default theme