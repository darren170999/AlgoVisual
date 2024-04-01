import React, { useEffect, useState } from "react";
import { ChakraProvider, Box, Text } from "@chakra-ui/react";
import Header from "../components/Header";
import HomeLoggedIn from "./HomeLoggedIn";
import IntroComponent from "../components/IntroComponent";
import Announcements from "../components/Announcements";

function Home() {
  const [isloggedIn, setIsLoggedIn] = useState<boolean>(true);
  const [displayText, setDisplayText] = useState<string>("");

  useEffect(() => {
    const check = localStorage.getItem("user");
    if(check === "true"){
      setIsLoggedIn(true);
    }

    if (isloggedIn!) {
      const text = "Welcome to algovisor";
      let index = 0;
      const intervalId = setInterval(() => {
        setDisplayText(text.slice(0, index));
        index++;
        if (index > text.length) {
          clearInterval(intervalId);
          setDisplayText(prev => prev + "|"); // Add cursor after completion
        }
      }, 200);
    }
  }, []);

  return (
    <ChakraProvider>
      <Header />
      <Box position="relative">
        {isloggedIn! ? (
          <>
          <IntroComponent/>
            <Box
            position="absolute"
            top="20%"
            left="30%"
            transform="translate(-50%, -50%)"
            textAlign="center"
            zIndex={2}
              >
                <Text fontSize="5xl" color="white" fontWeight="bold" fontFamily={"arial"}>
                  {displayText}
                </Text>
            </Box>
          </>
        ) : 
        <>
          <HomeLoggedIn/>
        </>}
        
      </Box>
    </ChakraProvider>
  );
}

export default Home;
