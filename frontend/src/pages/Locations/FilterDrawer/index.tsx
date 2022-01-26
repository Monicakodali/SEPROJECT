import { Drawer, Typography} from '@mui/material';
import { styled } from '@mui/material/styles';
import * as React from 'react';

const drawerWidth = 240

const StyledDrawer = styled(Drawer)(({ theme }) => ({
  // flexShrink: 0,
  '& .MuiDrawer-paper': {
    position: 'absolute',
    boxSizing: 'border-box',
    padding: theme.spacing(2),
  },
  '& .MuiDrawer-modal': {
    position: 'absolute',
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
    <Typography sx={{fontWeight: 'bold'}}>
      Filters
    </Typography>
  </StyledDrawer>
  );
};