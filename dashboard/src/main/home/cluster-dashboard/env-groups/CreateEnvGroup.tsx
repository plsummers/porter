import React, { Component } from "react";
import styled from "styled-components";

import sliders from "assets/sliders.svg";

import { Context } from "shared/Context";
import InputRow from "components/values-form/InputRow";
import KeyValueArray from "components/values-form/KeyValueArray";
import Selector from "components/Selector";
import Helper from "components/values-form/Helper";
import SaveButton from "components/SaveButton";
import { isAlphanumeric } from "shared/common";

type PropsType = {
  goBack: () => void;
};

type StateType = {
  expand: boolean;
  update: any[];
  envGroupName: string;
  selectedNamespace: string;
  namespaceOptions: any[];
  envVariables: any;
};

export default class CreateEnvGroup extends Component<PropsType, StateType> {
  state = {
    expand: false,
    update: [] as any[],
    envGroupName: "",
    selectedNamespace: "",
    namespaceOptions: [] as any[],
    envVariables: {} as any,
  };

  isDisabled = () => {
    return !(!isAlphanumeric(this.state.envGroupName) &&
    this.state.envGroupName !== "");
  }

  render() {
    return (
      <>
        <StyledCreateEnvGroup>
          <HeaderSection>
            <Button onClick={this.props.goBack}>
              <i className="material-icons">
                keyboard_backspace
              </i>
              Back
            </Button>
            <Title>Create an Environment Group</Title>
          </HeaderSection>
          <DarkMatter antiHeight="-13px" />
          <Heading isAtTop={true}>Name</Heading>
          <Subtitle>
            Randomly generated if left blank.
            <Warning
              highlight={
                !isAlphanumeric(this.state.envGroupName) &&
                this.state.envGroupName !== ""
              }
            >
              Lowercase letters, numbers, and "-" only.
            </Warning>
          </Subtitle>
          <DarkMatter antiHeight="-29px" />
          <InputRow
            type="text"
            value={this.state.envGroupName}
            setValue={(x: string) => this.setState({ envGroupName: x })}
            placeholder="ex: doctor-scientist"
            width="100%"
          />

          <Heading>Destination</Heading>
          <Subtitle>
            Specify the namespace you would like to create this environment group in.
          </Subtitle>
          <DestinationSection>
            <NamespaceLabel>
              <i className="material-icons">view_list</i>Namespace
            </NamespaceLabel>
            <Selector
              key={"namespace"}
              activeValue={this.state.selectedNamespace}
              setActiveValue={(namespace: string) =>
                this.setState({ selectedNamespace: namespace })
              }
              options={this.state.namespaceOptions}
              width="250px"
              dropdownWidth="335px"
              closeOverlay={true}
            />
          </DestinationSection>

          <Heading>Environment Variables</Heading>
          <Helper>
            Set environment variables for your secrets and environment-specific configuration.
          </Helper>
          <KeyValueArray
            values={this.state.envVariables}
            setValues={(x: any) => this.setState({ envVariables: x })}
          />
          <SaveButton
            disabled={this.isDisabled()}
            text="Deploy"
            onClick={() => console.log("asdf")}
            status={
              this.isDisabled()
                ? "Missing required fields"
                : "What is the status?"
            }
            makeFlush={true}
          />
        </StyledCreateEnvGroup>
        <Buffer />
      </>
    );
  }
}

CreateEnvGroup.contextType = Context;

const Buffer = styled.div`
  width: 100%;
  height: 150px;
`;

const StyledCreateEnvGroup = styled.div`
  padding-bottom: 50px;
  position: relative;
`;

const NamespaceLabel = styled.div`
  margin-right: 10px;
  display: flex;
  align-items: center;
  > i {
    font-size: 16px;
    margin-right: 6px;
  }
`;

const DestinationSection = styled.div`
  display: flex;
  align-items: center;
  color: #ffffff;
  font-family: "Work Sans", sans-serif;
  font-size: 14px;
  margin-top: 2px;
  font-weight: 500;
  margin-bottom: 32px;

  > i {
    font-size: 25px;
    color: #ffffff44;
    margin-right: 13px;
  }
`;

const Button = styled.div`
  display: flex;
  flex-direction: row;
  align-items: center;
  justify-content: space-between;
  font-size: 13px;
  cursor: pointer;
  font-family: "Work Sans", sans-serif;
  border-radius: 20px;
  color: white;
  height: 35px;
  margin-left: -2px;
  padding: 0px 8px;
  padding-bottom: 1px;
  font-weight: 500;
  padding-right: 15px;
  overflow: hidden;
  white-space: nowrap;
  text-overflow: ellipsis;
  cursor: pointer;
  border: 2px solid #969fbbaa;
  :hover {
    background: #ffffff11;
  }

  > i {
    color: white;
    width: 18px;
    height: 18px;
    color: #969fbbaa;
    font-weight: 600;
    font-size: 14px;
    border-radius: 20px;
    display: flex;
    align-items: center;
    margin-right: 5px;
    justify-content: center;
  }
`;

const DarkMatter = styled.div<{ antiHeight?: string }>`
  width: 100%;
  margin-top: ${(props) => props.antiHeight || "-15px"};
`;

const Warning = styled.span<{ highlight: boolean; makeFlush?: boolean }>`
  color: ${(props) => (props.highlight ? "#f5cb42" : "")};
  margin-left: ${(props) => (props.makeFlush ? "" : "5px")};
`;

const Subtitle = styled.div`
  padding: 11px 0px 16px;
  font-family: "Work Sans", sans-serif;
  font-size: 13px;
  color: #aaaabb;
  line-height: 1.6em;
  display: flex;
  align-items: center;
`;

const Title = styled.div`
  font-size: 24px;
  font-weight: 600;
  font-family: "Work Sans", sans-serif;
  margin-left: 15px;
  border-radius: 2px;
  color: #ffffff;
`;

const HeaderSection = styled.div`
  display: flex;
  align-items: center;
  margin-bottom: 40px;

  > i {
    cursor: pointer;
    font-size 20px;
    color: #969Fbbaa;
    padding: 2px;
    border: 2px solid #969fbbaa;
    border-radius: 100px;
    :hover {
      background: #ffffff11;
    }
  }

  > img {
    width: 20px;
    margin-left: 17px;
    margin-right: 7px;
  }
`;

const Heading = styled.div<{ isAtTop?: boolean }>`
  color: white;
  font-weight: 500;
  font-size: 16px;
  margin-bottom: 5px;
  margin-top: ${(props) => (props.isAtTop ? "10px" : "30px")};
  display: flex;
  align-items: center;
`;