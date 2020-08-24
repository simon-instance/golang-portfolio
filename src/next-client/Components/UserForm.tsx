import React, { useState, useEffect } from "react";
import {
    TryCatchHandler,
    TryCatchProps,
    TryCatchDataProps,
    RequestHandler,
    RequestProps,
} from "../Handlers";
import {
    Box,
    Heading,
    FormControl,
    FormLabel,
    Input,
    Button,
    useToast,
} from "@chakra-ui/core";
import Wrapper from "./partials/Wrapper";
import { Formik, setNestedObjectValues, FormikProps } from "formik";
import InputField from "./partials/forms/InputField";

//
// LoginHeader: text to inform user about what to do
//

const LoginHeader: React.FC<{ readonly type: string }> = ({ type }) => {
    return (
        <Box>
            <Heading as="h3">
                {type === "login" ? "Log in" : "Registreer"} met uw gegevens
            </Heading>
        </Box>
    );
};

//
// Login: main component
//

const UserForm: React.FC<{ type: string }> = ({ type }) => {
    const [username, setUsername] = useState<string | null>(null);
    const [password, setPassword] = useState<string | null>(null);

    const TCHandler: TryCatchProps = new TryCatchHandler();
    const RHandler: RequestProps = new RequestHandler();
    const toast = useToast();

    const HandleSubmit: () => Promise<void> = async (
        values,
        { setErrors, setStatus }
    ) => {
        if (username !== null && password !== null) {
            RHandler.data = {
                Username: username,
                Password: password,
            };
            RHandler.method = "POST";
            RHandler.uri = `/api/auth/${type === "login" ? type : "register"}`;

            // let data: TryCatchProps = TCHandler.Data;
            // let toastConfig = {
            //     title: null,
            //     description: null,
            // };

            // try {
            //     const response: any = await fetch(
            //         `http://127.0.0.1:8080/api/auth/${
            //             type === "login" ? type : "register"
            //         }`,
            //         requestOptions
            //     );

            //     data.response.data = await response.json();
            //     data.response.status = response.status;

            //     const error: Error | void = TCHandler.handleData(data);
            //     if (typeof error !== "undefined") throw error;

            //     toastConfig = {
            //         title: "Gelukt",
            //         description: `Je bent nu ${
            //             type === "login" ? "ingelogd" : "geregistreed"
            //         }`,
            //     };
            // } catch (e) {
            //     data.error.status = true;
            //     data.error.message = e;
            //     toastConfig = {
            //         title: "Foutmelding",
            //         description: e.message,
            //     };
            // } finally {
            //     console.log(data.error.status);
            //     TCHandler.Data = data;
            //     toast({
            //         position: "top-left",
            //         status: data.error.status === true ? "error" : "success",
            //         duration: 7000,
            //         ...toastConfig,
            //         isClosable: true,
            //     });
            // }
        }
    };

    return (
        <Wrapper>
            <Formik
                initialValues={{ username: "", password: "" }}
                onSubmit={HandleSubmit}
            >
                {(
                    props: FormikProps<{ username: string; password: string }>
                ) => (
                    <form onSubmit={props.handleSubmit}>
                        <InputField
                            name="username"
                            type="text"
                            placeholder="Username"
                            label="Username"
                        />
                        <Box mt={5}>
                            <InputField
                                name="password"
                                type="password"
                                placeholder="Password"
                                label="Password"
                            />
                        </Box>
                        <Button
                            w="100%"
                            mt={5}
                            variantColor="teal"
                            type="submit"
                            isLoading={props.isSubmitting}
                        >
                            Login
                        </Button>
                    </form>
                )}
            </Formik>
        </Wrapper>
    );
};

export default UserForm;
