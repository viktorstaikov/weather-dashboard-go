import Container from '@material-ui/core/Container';
import { withStyles } from '@material-ui/core/styles';
import Typography from '@material-ui/core/Typography';
import PropTypes from 'prop-types';
import React, { Component } from 'react';
import { Cell, Pie, PieChart } from 'recharts';

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
export class UvChart extends Component {
  render() {
    const { uvIndex, classes } = this.props;

    const data = [{ value: uvIndex }, { value: 12 - uvIndex }];

    return (
      <Container className={classes.root}>
        <div className={classes.chart}>
          <PieChart width={200} height={100}>
            <Pie dataKey="value" startAngle={180} endAngle={0} data={data} innerRadius={60} outerRadius={80} cy={85}>
              <Cell fill="blue" />
              <Cell fill="gray" />
            </Pie>
          </PieChart>
        </div>
        <div className={classes.legend}>
          <Typography variant="h6" gutterBottom>
            UV Index: {uvIndex}
          </Typography>
        </div>
      </Container>
    );
  }
}

UvChart.propTypes = {
  speed: PropTypes.number.isRequired,
  deg: PropTypes.number.isRequired,
};

export default withStyles(styles)(UvChart);
