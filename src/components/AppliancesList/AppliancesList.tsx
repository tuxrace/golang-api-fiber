import React, { useEffect, useState } from "react";
import axios from "axios";
import { API_URL } from "../../config";
import {
  Card,
  CardHeader,
  makeStyles,
  Grid,
  CardContent,
  Box,
} from "@material-ui/core";
import Search from "../Search";
import styles from "./Appliances.styles";

const useStyles = makeStyles(styles);

const AppliancesList = () => {
  const [appliances, setAppliances] = useState([]);
  const classes = useStyles();

  useEffect(() => {
    async function getAppliances() {
      const res = await axios.get(`${API_URL}/appliances`);
      console.log(res);
      setAppliances(res.data);
    }
    getAppliances();
  }, []);

  return (
    <Grid container direction="row" spacing={2} alignItems="center">
      <Grid item>
        <h1>Appliances</h1>
      </Grid>
      <Grid item xs={12} lg={12}>
        <Search />
      </Grid>
      <Grid item container direction="column" spacing={2}>
        {appliances.map((appliance: any) => (
          <Box py={1}>
            <Card className={classes.card} raised>
              <CardHeader title={appliance.model} />
              <CardContent>{appliance.model}</CardContent>
            </Card>
          </Box>
        ))}
      </Grid>
    </Grid>
  );
};

export default AppliancesList;
