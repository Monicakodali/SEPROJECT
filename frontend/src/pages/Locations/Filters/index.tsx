import { Box, Typography} from '@mui/material';
import * as React from 'react';

type FilterDrawerProps = {
}

export default function FilterDrawer() {
  
  return (<Box sx={{margin: 2}}>
      <Typography sx={{fontWeight: 'bold'}}>
        Filters
      </Typography>
    </Box>
  );
};