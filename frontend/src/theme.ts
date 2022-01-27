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
});

export default theme