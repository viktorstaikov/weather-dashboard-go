import CssBaseline from '@material-ui/core/CssBaseline';
import { withStyles } from '@material-ui/core/styles';

import PropTypes from 'prop-types';
import React, { Component } from 'react';
import './App.css';
import Header from './components/Header';
import Dashboard from './pages/Dashboard';

const styles = theme => ({
  root: {
    display: 'flex'
  },
  appBarSpacer: theme.mixins.toolbar,
  content: {
    flexGrow: 1,
    height: '100vh',
    overflow: 'auto'
  }
});
export class App extends Component {
  render() {
    const { classes } = this.props;

    return (
      <div className={classes.root}>
        <CssBaseline />

        <Header></Header>

        <main className={classes.content}>
          <div className={classes.appBarSpacer} />

          <Dashboard></Dashboard>
        </main>
      </div>
    );
  }
}

App.propTypes = {
  classes: PropTypes.object.isRequired
};

export default withStyles(styles)(App);
