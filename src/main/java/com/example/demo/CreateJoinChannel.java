package com.example.demo;

import org.hyperledger.fabric.sdk.*;

import java.io.File;
import java.util.Collection;
import java.util.LinkedList;
import java.util.Properties;
import java.util.concurrent.TimeUnit;

import static org.hyperledger.fabric.sdk.Channel.PeerOptions.createPeerOptions;

/**
 * Created by Mujji on 3/4/2018.
 */
public class CreateJoinChannel {

    Channel constructChannel(String name, HFClient client, SampleOrg sampleOrg, Channel newChannel) throws Exception {
        ////////////////////////////
        //Construct the channel
        //

        System.out.println("Joining channel "+ name);

        //boolean doPeerEventing = false;
//		boolean doPeerEventing = !testConfig.isRunningAgainstFabric10() && BAR_CHANNEL_NAME.equals(name);
//        boolean doPeerEventing = !testConfig.isRunningAgainstFabric10() && FOO_CHANNEL_NAME.equals(name);
        //Only peer Admin org
        client.setUserContext(sampleOrg.getPeerAdmin());


        boolean everyother = true; //test with both cases when doing peer eventing.
        for (String peerName : sampleOrg.getPeerNames()) {
            String peerLocation = sampleOrg.getPeerLocation(peerName);


            newChannel.initialize();

            System.out.println("Finished initialization channel "+ name);

            //Just checks if channel can be serialized and deserialized .. otherwise this is just a waste :)
            byte[] serializedChannelBytes = newChannel.serializeChannel();
            newChannel.shutdown(true);

            return client.deSerializeChannel(serializedChannelBytes).initialize();

        }
        return newChannel;
    }
}
