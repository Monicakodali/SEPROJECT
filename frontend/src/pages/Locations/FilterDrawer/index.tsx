import { Drawer, Toolbar, Typography} from '@mui/material';
import { styled } from '@mui/material/styles';
import * as React from 'react';

const drawerWidth = 240

const StyledDrawer = styled(Drawer)(({ theme }) => ({
  // flexShrink: 0,
  '& .MuiDrawer-paper': {
    boxSizing: 'border-box',
    padding: theme.spacing(2),
    zIndex: 0,
  },

}));

type FilterDrawerProps = {
  drawerWidth: number
}

export default function FilterDrawer({drawerWidth}: FilterDrawerProps) {
  
  return (<StyledDrawer
      open
      variant="persistent"
      anchor="left"
      sx={{
        '& .MuiDrawer-paper': {
          width: drawerWidth,
        }
      }}
    >
    <Toolbar />
    <Typography sx={{fontWeight: 'bold'}}>
      Filters
    </Typography>
  </StyledDrawer>
  );
};