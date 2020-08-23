import React, { InputHTMLAttributes, PropsWithChildren } from "react";
import {
    FormControl,
    FormLabel,
    Input,
    FormErrorMessage,
} from "@chakra-ui/core";

import { useField } from "formik";

type InputFieldProps = InputHTMLAttributes<HTMLInputElement> & {
    name: string;
    label: string;
};

const InputField: React.FC<InputFieldProps> = (props) => {
    const [field, { error, status }] = useField(props);

    return (
        <FormControl isInvalid={!!error}>
            <FormLabel htmlFor={field.name}>{props.label}</FormLabel>
            <Input {...field} id={field.name} />
            <FormErrorMessage>{}</FormErrorMessage>
        </FormControl>
    );
};

export default InputField;
