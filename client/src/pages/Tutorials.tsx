import FullScreenSection from "../components/FullScreenSection";
import { Outlet } from 'react-router-dom';
import Header from "../components/Header";
import { Text, Avatar, Button, ButtonGroup, Card, CardBody, CardFooter, Divider, Heading, Stack, Tab, TabList, TabPanel, TabPanels, Tabs, VStack, HStack, Box } from "@chakra-ui/react";
import AccordionRows from "../components/AccordionRows";
import TutorialQuestion from "../components/TutorialQuestion";
import { useEffect, useState } from "react";
import { QnType } from "../types/QnType";
import { saveAttemptDataProps } from "../types/SaveAttemptDataProps";
function Tutorials(){
  //PUT IN DB, Call everything in Tutorials one time and filter. Status is all new. 
  const arrayList = ["Arrays & Hashing", "Two Pointers", "Stack", "Binary Search", "Sliding Window", "Linked List", 
  "Trees", "Tries", "Backtracking", "Heaps", "Graphs", "Dynamic Programming"]
  const [questions, setQuestions] = useState<QnType[]>([]);
  const [completedQns, setCompletedQns] = useState<string[]>([]);
  let username = localStorage.getItem("username");
  const getAttempts = async() =>{
    try {
      const response = await fetch(`https://algovisor.onrender.com/tutorials/code/check/${username}`)
      if(response.ok){
        const res = await response.json();
        const result: saveAttemptDataProps[] = res.data.data;
        const completedQnsIds = result.filter(attempt => attempt.status === "Completed")
        .map(completedAttempt => completedAttempt.qnid);
        setCompletedQns(completedQnsIds);

      }
    } catch (error) {
    }
  }
  useEffect(()=>{
    getAttempts();

  },[])

  const getQuestions = async() => {
    try{
        const response = await fetch("https://algovisor.onrender.com/tutorials", {
            method: "GET",
            headers: {
                "Content-Type": "application/json",
            }
        });
        if(response.ok){
            var res = await response.json();
            setQuestions(res.data.data)
            
        }
    } catch {
    }
  }
  useEffect(()=>{
      getQuestions()
  },[])

  const appropriateTabs = (tabIndex: number) => {
    const tabName = `Tut${tabIndex + 1}`;
    return questions.filter((question) => question.tags === tabName)
    .map((question)=> {
      const boolCompleted = completedQns.includes(question.qnid);
      return(
        <TutorialQuestion
          key={question.qnid}
          qnid={question.qnid}
          status={question.status}
          tags={question.tags}
          description={question.description}
          name={question.name}
          boolCompleted={boolCompleted} 
          examples={""} constraints={""}/>
      );
    });
  }

  return(
      <>
      <Header></Header>
      <FullScreenSection backgroundColor="#1a1f71" alignItems="center" flex-direction="row"
      isDarkBackground p={8}  spacing={8}>
      <VStack>
        {" "}
        <h1 style={{color: "#F7B600"}}></h1>
        <Heading>
          {" "}
          <Box textAlign="center" mb={4}>
            <Heading as="h2" size="lg" color="white">Algo Code</Heading>
          </Box>
          <Box textAlign="center" mb={4}>
            <Heading as="h2" size="md" color="white">Practice algorithms segregated by topics</Heading>
          </Box>
          {" "}
        </Heading>
      </VStack>
      <VStack>
      <Tabs variant='enclosed' color={"yellow"} maxWidth=''>
        <TabList>
          {[...Array(12).keys()].map((index) => (
            <Tab key={index}>AC {index + 1}</Tab>
          ))}
        </TabList>
        <TabPanels>
          {[...Array(12).keys()].map((index) => (
                <TabPanel key={index} className={`Tut${index + 1}`}>
                  <Box overflowX="auto" w="100%" maxWidth="900px">
                    <Box textAlign="left" mb={4} padding={"10px"}>
                      <Heading as="h2" size="sm" color="white">{arrayList[index]}</Heading>
                    </Box>
                    <HStack spacing={4} style={{ flexWrap: "nowrap" }}>
                      {appropriateTabs(index)}
                    </HStack>
                  </Box>
                </TabPanel>
              ))}
        </TabPanels>
      </Tabs>
      </VStack>
      </FullScreenSection>
      </>
  );
}
export default Tutorials;
