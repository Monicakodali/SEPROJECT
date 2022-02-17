import * as React from 'react';
import { Box, SxProps, Typography } from '@mui/material';
import { styled } from '@mui/material/styles';

type RatingProps = {
  rating: number,
  numRatings?: number,
  size: number,
  label?: string,
  labelStyles?: SxProps
}

const Star = styled('div', {
  shouldForwardProp: (prop) => prop !== "size" && prop !=="active"
})<{size: number, active: boolean}>(({theme, size, active}) => ({
  display: 'inline-flex',
  alignItems: 'center',
  justifyContent: 'center',
  backgroundColor: !active ? 'lightgray' : 'rgb(255,126,66)',
  lineHeight: 1,
  fontSize: size*0.85,
  borderRadius: size/4,
  color: '#FFF',
  height: size,
  width: size,
  '&:not(:last-child)': {
    marginRight: size/4
  }
}))




const star = '★'
const MAX_RATING = 5
const chars = (new Array(MAX_RATING)).fill(star)

export default function Rating({rating, numRatings, label, size, labelStyles}: RatingProps) {
  
  return (
    <Box sx={{display: 'flex', alignItems: 'center'}}>
      <Box sx={{mr: 1.5}}>
        {chars.map((c, i) => <Star key={i} active={(i+1) <= rating} size={size}>{c}</Star>)}
      </Box>
      <Typography color="text.secondary" sx={{...(labelStyles || {}), fontSize: size*0.6}}>
        {numRatings}{label ? ` ${label}` : ''}
      </Typography>
    </Box>
  );
};

Rating.defaultProps = {
  size: 16
}