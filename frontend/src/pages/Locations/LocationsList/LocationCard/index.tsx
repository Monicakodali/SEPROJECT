import * as React from 'react';
import { Box, CardContent, Typography, CardMedia, ListItemButton } from '@mui/material';
import { styled } from '@mui/material/styles';

import { Location } from '../../../../types/location'

type LocationsCardProps = {
  data: Location,
  selected?: boolean,
  onClick: () => void
}

const placeholderImageUrl = 'https://thelivingstonpost.com/wp-content/themes/fox/images/placeholder.jpg'

const StyledListItem = styled(ListItemButton)(({ theme }) => {

  return {
    display: 'flex',
    borderWidth: 1,
    borderColor: theme.palette.grey[300],
    borderStyle: 'solid',
    padding: 0,
    '&:not(:last-child)': {
      marginBottom: theme.spacing(2)
    },
    '&:hover, &:focus': {
      boxShadow: theme.shadows[3],
      //backgroundColor: 'initial' //modify hover bg color?
    }
  }
})

export default function LocationsCard({data, selected, onClick}: LocationsCardProps) {

  const { name, id } = data
  
  return (
    <StyledListItem selected={selected} alignItems="flex-start" focusRipple={false} onClick={onClick}>
      <CardMedia
        component="img"
        sx={{ width: 151, padding: 2 }}
        image={placeholderImageUrl}
        alt={name + ' image'}
      />
      <Box sx={{ display: 'flex', flexGrow: 1, flexDirection: 'column' }}>
        <CardContent sx={{ flex: '1 0 auto', px: 1, py: 3 }}>
          <Typography component="div" variant="h5">
            {name}
          </Typography>
          <Typography variant="subtitle1" color="text.secondary" component="div">
            {id}
          </Typography>
        </CardContent>
      </Box>

    </StyledListItem>
  );
};