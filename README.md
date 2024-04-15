- CL-App is a command line utility for converting CSV to JSON file formats. Currently this exe is only formated to convert a specific format of CSV utilizing both integer and float datatypes.\
- Data output is formatted with each line being a different item.\
- The first field is capable of handling scientific notation and all values are assumed to be positive values.

- Unit testing was utilized on the 'ParseLine' and 'ValidateArgs' functions since they are the primary breaking points within the program
- Concurrency has not been implemented due to time constraints and time being spent towards attempting a general CSV to JSON solution for all data types and number of fields.
--------USAGE-------------
./CSVtoJSON.exe <InputFile>.csv <OutputFile>.json

additional help information can be found by running ./CSVtoJSON.exe without any arguments
