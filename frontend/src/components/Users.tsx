import React, { useEffect } from "react";

import { Link as RouterLink } from "react-router-dom";

import Typography from "@mui/material/Typography";

import Button from "@mui/material/Button";

import Container from "@mui/material/Container";

import Box from "@mui/material/Box";


import { DataGrid, GridColDef } from "@mui/x-data-grid";

import { createTheme, ThemeProvider } from "@mui/material/styles";
import { green } from "@mui/material/colors";
import { ScheduleInterface } from "../models/ISchedule";
import { Table, TableBody, TableCell, TableContainer, TableHead, TableRow } from "@mui/material";
import moment from "moment";
import { DoctorInterface } from "../models/IDoctor";
// import { createStyles, makeStyles,Theme } from '@mui/material/styles';
// import { makeStyles,Theme,createStyles } from "@mui/material";
const theme = createTheme({
  palette: {
    primary: {
      // Purple and green play nicely together.
      main: green[500],
    },
    secondary: {
      // This is green.A700 as hex.
      main: "#e8f5e9",
    },
  }
});


// const useStyles = makeStyles((theme: Theme) =>

//  createStyles({

//    container: {marginTop: theme.spacing(2)},

//    table: { minWidth: 650},

//    tableSpace: {marginTop: 20},

//  })

// );


function Users() {
  // const classes = useStyles();
  const [schedule, setSchedule] = React.useState<ScheduleInterface[]>([]);
  const [user, setUser] = React.useState<DoctorInterface>();

  const getSchedule = async () => {
    const apiUrl = "http://localhost:8080/schedules";

    const requestOptions = {
      method: "GET",

      headers: { Authorization: `Bearer ${localStorage.getItem("token")}`, "Content-Type": "application/json" },
    };

    fetch(apiUrl, requestOptions)
      .then((response) => response.json())

      .then((res) => {
        console.log(res.data);

        if (res.data) {
          setSchedule(res.data);
        }
      });
  };
console.log(schedule)
  // //ตาราง
  // const columns: GridColDef[] = [
  //   { field: "ID", headerName: "ID", width: 50 },

  //   { field: "Name", headerName: "Name", width: 200 },

  //   { field: "Medical activity", headerName: "Medical activity", width: 250 },

  //   { field: "WorkPlace", headerName: "Location", width: 250 },

  //   { field: "Time", headerName: "Date & Time", width: 200 },
  // ];

  useEffect(() => {
    const getToken = localStorage.getItem("token");
        if (getToken) {
            setUser(JSON.parse(localStorage.getItem("user") || ""));
        }
    getSchedule();
  }, []);

  return (
    <ThemeProvider theme={theme}>
      <div>
        <Container maxWidth="md">
          <Box
            display="flex"
            sx={{
              marginTop: 2,
            }}
          >
            <Box flexGrow={1}>
              <Typography // ตารางเวลา
                component="h1"
                variant="h6"
                color="inherit"
                gutterBottom
              >
                ตารางเวลาแพทย์
              </Typography>
            </Box>

            <Box>
              <Button //ตัวบันทึก
                component={RouterLink} //ลิ้งหน้าต่อไป
                to="/Schedule"
                variant="contained"
                color="primary"
              >
                <Typography
                  color="secondary"
                  component="div"
                  sx={{ flexGrow: 1 }}
                >
                  บันทึกข้อมูลตารางเวลาแพทย์
                </Typography>
              </Button>
            </Box>
          </Box>

          <div style={{ height: 400, width: "100%", marginTop: "20px" }}>
            <TableContainer >
              <Table  aria-label="simple table">
                <TableHead> 
                  {/* หัวข้อตาราง */}
                  <TableRow>
                    <TableCell align="left" width="15%">
                      ID
                    </TableCell>

                    <TableCell align="center" width="15%">
                    Name
                    </TableCell>

                    <TableCell align="center" width="15%">
                    activity
                    </TableCell>

                    <TableCell align="center" width="20%">
                    WorkPlace
                    </TableCell>
                    <TableCell align="center" width="20%">
                    Time
                    </TableCell>
                  </TableRow>
                </TableHead>
{/* ดึงช้อมูล */}
                <TableBody>
                  {schedule.map((item: ScheduleInterface) => (
                    <TableRow key={item.ID}>
                      <TableCell align="left">{item.ID}</TableCell>

                      <TableCell align="left">{item.Doctor?.Name}
                       
                      </TableCell>

                      <TableCell align="left">{item.MedActivity?.Name}</TableCell>

                      <TableCell align="center">{item.WorkPlace?.Name}</TableCell>

                      <TableCell align="center">     
                        {moment(item.Time).format("DD/MM/YYYY HH:mm:ss A")}
                      </TableCell>
                    </TableRow>
                  ))}
                </TableBody>
              </Table>
            </TableContainer>

            {/* <DataGrid

           rows={schedule}

           getRowId={(row) => row.ID}

           columns={columns}

           pageSize={5}

           rowsPerPageOptions={[5]}

         /> */}
          </div>
        </Container>
      </div>
  </ThemeProvider>
  );
}

export default Users;

