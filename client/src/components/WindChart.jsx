import Arrow from '@elsdoerfer/react-arrow';
import Container from '@material-ui/core/Container';
import { withStyles } from '@material-ui/core/styles';
import Typography from '@material-ui/core/Typography';
import PropTypes from 'prop-types';
import React, { Component } from 'react';

const styles = theme => ({
  root: {
    display: 'flex',
    flexFlow: 'row nowrap',
  },
  chart: {
    flex: '1 0 200px',
  },
  legend: {
    flex: '1 1 100px',
  },
});

const directions = ['North', 'NE', 'East', 'SE', 'South', 'SW', 'West', 'NW'];

export class WindChart extends Component {
  render() {
    const { classes, speed, deg } = this.props;

    const slice = 360 / directions.length;
    const adjusted = deg - slice / 2;
    const idx = Math.ceil(adjusted / slice);
    const dir = directions[idx];
    return (
      <Container className={classes.root}>
        <div className={classes.chart}>
          <Arrow
            angle={deg}
            length={80}
            lineWidth={speed}
            style={{
              width: '100px',
              height: '100%',
            }}
          />
        </div>
        <div className={classes.legend}>
          <Typography variant="h6" gutterBottom>
            {Math.floor(speed)} m/sec
          </Typography>
          <Typography variant="subtitle1" gutterBottom>
            {dir}
          </Typography>
        </div>
      </Container>
    );
  }
}

WindChart.propTypes = {
  speed: PropTypes.number.isRequired,
  deg: PropTypes.number.isRequired,
};

export default withStyles(styles)(WindChart);
