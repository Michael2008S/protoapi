// Code generated by protoapi; DO NOT EDIT.

package com.yoozoo.spring;

import com.fasterxml.jackson.annotation.JsonCreator;
import com.fasterxml.jackson.annotation.JsonProperty;

import java.util.List;

public class Tag {
    private final int tag_id;
    private final String tag_name;

    @JsonCreator
    public Tag(@JsonProperty("tag_id") int tag_id, @JsonProperty("tag_name") String tag_name) {
        this.tag_id = tag_id;
        this.tag_name = tag_name;
    }

    public int getTag_id() {
        return tag_id;
    }
    public String getTag_name() {
        return tag_name;
    }
    
}
